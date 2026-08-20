# Change: add-stablecoin-token-type

> **Status:** planning. Contracts work is the blocking foundation; ops-api + Blockscout edits are inert until it lands. The token-type classification fix this builds on has **already landed separately**.

---

## 1. Why

Product wants to deploy Circle-style stablecoins (per [circlefin/stablecoin-evm](https://github.com/circlefin/stablecoin-evm), `FiatTokenV2`) through the Rayls factory and have them surface in ops-api as a distinct token type, `rayls-stable-coin` — the same way `RAYLS_ERC20` is a distinct type today.

The naive idea — "just register Circle's `FiatTokenV2` runtime bytecode with `setBytecode` and deploy it" — **does not work** with this stack. Neither does "inherit Circle's contracts alongside ours." See [§3, Why not integrate Circle directly](#3-why-not-integrate-circle-directly) for the full five-reason analysis. In short: the factory forces a fixed Rayls init dispatch Circle can't satisfy, Circle is a proxy/delegatecall design while our factory does raw-runtime CREATE2, the two storage layouts collide, and Blockscout detection needs a Rayls event Circle never emits.

The chosen approach is a **Rayls-native `RaylsStableCoinHandler`** — a contract that ports the stablecoin *feature surface* (pause, blacklist, master-minter/controlled minting) onto the existing audited `RaylsErc20Handler` base, emits a new `RaylsStableCoinCreated` event, and is deployed under a new factory key `RAYLS_STABLECOIN`. ops-api deploys it via the **already-generated** `deployRegistered` binding (no abigen regeneration) and detects it as `rayls-stable-coin`.

This change spans three repos. The contracts work is the blocking foundation: nothing downstream deploys or is detected until the new handler + key + seeding land on-chain.

---

## 2. Context — how a "token type" works here

A token "type" in this stack is not a single artifact — it is four coordinated pieces, all keyed off a `bytes32` factory key and a creation event:

1. **Handler contract** — implements `IRaylsInitializer.initialize(bytes userArgs, RaylsTrustedInit trusted)` and emits a creation event. The factory calls this `initialize` after CREATE2.
2. **Concrete production contract** — compiled to non-empty `deployedBytecode`, whose runtime is seeded into the factory via `setBytecode(KEY, runtime)` at PN deploy time.
3. **Blockscout discovery** — maps the creation event's topic hash → a `tokens.type` string.
4. **ops-api mapping** — maps that string → `ErcStandard` enum → `Label()` for API/SSE responses.

This is exactly how `RAYLS_ERC20` works today. We add a new standard, `RAYLS_STABLECOIN`, by replicating each piece.

---

## 3. Why not integrate Circle directly

Five concrete, independent blockers (each verified against Circle's source and the Rayls factory). Any one alone rules out "use the bytecode"; together they make even "fork their Solidity into our SDK" a large, risky port rather than a phase-1 task.

**3.1 — The factory's init dispatch is fixed, and Circle's `initialize` doesn't match it.**
The factory deploys, then unconditionally calls one selector (`AbstractContractFactoryV1.sol:320`):
```solidity
deployed.call(abi.encodeCall(IRaylsInitializer.initialize, (userArgs, trusted)))
```
Selector: `initialize(bytes,(address,address,address,address,bytes32,address))`. Circle's actual signature is `initialize(string,string,string,uint8,address,address,address,address)` — different selector, different ABI. The factory calls a selector stock Circle doesn't implement → revert (`FactoryV1__InitializationFailed`). No `userArgs` value can fix a wrong selector. **Stock bytecode physically cannot be deployed by this factory.**

**3.2 — Circle expects a proxy; the factory does raw-runtime CREATE2.**
`InitCodeStub` wraps a *finished runtime* and CREATE2-deploys it byte-for-byte as a standalone contract — no constructor execution, no proxy. Circle's `FiatTokenV1`/`V2` have **no constructor** and are designed to live **behind a `FiatTokenProxy`**, initialized via delegatecall (`initializeV#` per version). The two deployment architectures don't meet.

**3.3 — The init-guard mechanisms collide.**
Rayls uses OZ `Initializable` (`initialize(...) initializer`). Circle V1 uses a custom `bool initialized` guard; Circle V2 layers a version-gated reinitializer (`require(initialized && _initializedVersion == 0)`, then bumps `_initializedVersion`). Reconciling OZ's `_initialized` machinery with Circle's `initialized` bool **and** the multi-version `initializeV#` chain — all driven from a single factory `initialize` call — is a rewrite of Circle's initialization, and getting it subtly wrong bricks or re-initializes the instance.

**3.4 — Storage-layout collision (the silent, dangerous one).**
`RaylsErc20Handler` already declares fixed sequential slots (`tokenName`, `tokenSymbol`, `isCustom`, `internalDecimals`, `_erc20Identifier`, `lockedAmount`) plus inherited OZ `ERC20` (`_balances`, `_allowances`, `_totalSupply`, …). Circle declares its own from slot 0 (`name`, `symbol`, `decimals`, `currency`, `masterMinter`, `initialized`, `balanceAndBlacklistStates`, `allowed`, `totalSupply_`, `minters`, `minterAllowed`) — note Circle packs balance + blacklist state into **one** slot. You cannot multiple-inherit `RaylsErc20Handler` (OZ-ERC20-based) and Circle's `AbstractFiatTokenV1` (its own ledger): two overlapping balance ledgers at conflicting slots. The compiler won't always catch this; it corrupts balances at runtime.

**3.5 — Blockscout needs a Rayls event Circle never emits.**
Detection is event-based: the fetcher scans for a creation-event topic and writes `tokens.type`. Stock Circle emits no `Rayls*Created` event → even if it deployed, it would surface as plain `ERC-20`, never `rayls-stable-coin`. You'd have to modify Circle's code to emit it anyway.

**What we keep from Circle:** the *design* (the FiatToken role model + compliance controls), not the bytecode. Re-implementing those semantics on our audited base is strictly less work and less risk than reconciling two proxy models, two init guards, and two storage ledgers.

**One caveat to state plainly:** re-implementing means the result is **not** Circle's audited contract — you take on your own audit. If staying on Circle's audited code is a hard requirement, the only way to honor it is to run Circle's stack *outside* this factory (their proxy, their deploy) and bridge/wrap it — a much larger, different architecture decision, out of scope here.

---

## 4. What changes (by repo)

> Planning summary. ops-api is the capability owner (`tokens`); the contracts and Blockscout edits are external prerequisites tracked in [§7 Tasks](#7-tasks).

### rayls-privacy-contracts — blocking prerequisite, must land first
- Add `RAYLS_STABLECOIN_KEY = keccak256("RAYLS_STABLECOIN")` to `FactoryKeys.sol`; expose the constant on `AbstractContractFactoryV1`.
- Optionally add a typed `deployStableCoin(name, symbol, decimals, resourceId)` wrapper over `_deployRegistered` (functionally equal to generic `deployRegistered`).
- Add `RaylsStableCoinHandler.sol` under `src/rayls-protocol-sdk/tokens/`, extending `RaylsErc20Handler`, with the phase-1 surface (pause / blacklist / master-minter-controlled mint-burn), emitting `RaylsStableCoinCreated(address)` from `initialize`.
- Add `ProductionStableCoin.sol` under `src/rayls-protocol/prod-example-contracts/` (thin subclass for non-empty `deployedBytecode`, mirroring `ProductionErc20Token.sol`).
- Seed the key at PN deploy: add `{ label: 'StableCoin', keyFn: 'RAYLS_STABLECOIN_KEY', artifact: 'ProductionStableCoin' }` to the standards table in `hardhat/tasks/deploy/privacy-node.ts`.

### rayls-privacy-blockscout — prerequisite for detection
- Add `{"RaylsStableCoinCreated(address)", "Rayls-StableCoin"}` to `@event_signatures` in `rayls_token_discovery.ex`, and generalize the hardcoded 6-topic pin list (`[t1..t6]`) to handle the 7th signature.

### rayls-sovereign-ops-api — this repo
- Add `ErcStandardStableCoin ErcStandard = 7` to `domain/types.go`; `Label()` → `"RAYLS_STABLECOIN"`.
- Map `"Rayls-StableCoin"` → `ErcStandardStableCoin` in the shared `domain.ParseErcStandard` (`domain/types.go`) — one new `case`. Both the indexer and the repo filter already route through this function (see [§5 Dependencies](#5-dependencies)).
- Add a deploy case in `RaylsContractFactoryService.deployCalldata` (`cmd/api/adapters/blockchain/rayls_contract_factory.go`) → `PackDeployRegistered(keccak256("RAYLS_STABLECOIN"), abi.encode(name, symbol, decimals), bytes32(0))` (or `PackDeployStableCoin` if the typed wrapper is added).
- Accept the new standard in the deploy handler (`parseErcStandard` / `supportedStandards` in `cmd/api/adapters/handlers/token_deploy_handler.go`).
- Expose `ercStandardLabel = "RAYLS_STABLECOIN"` in responses — falls out automatically from `Label()`.

### Capabilities
- **Modified — `tokens`:** add `RAYLS_STABLECOIN` as a recognized ERC standard — deployable via `POST /api/tokens`, classified by the indexer, exposed in list/detail/SSE responses.

---

## 5. Dependencies

- **Token-type classification fix — ALREADY LANDED (separate change).** While scoping this we found a pre-existing bug: ops-api's `parseErcStandard`/`ercStandardFromString` matched only the plain Blockscout strings (`"ERC-20"`), but the `RaylsTokenDiscovery` fetcher writes Rayls-prefixed strings (`"Rayls-ERC-20"`, …) into `tokens.type`, so every Rayls-native token classified as `CUSTOM`. Fixed by introducing `domain.ParseErcStandard` (handling all three string forms — Rayls-prefixed, plain ERC, and canonical labels) and routing both call sites through it. This change builds on that shared function — adding `RAYLS_STABLECOIN` is now a single new `case`.

---

## 6. Decisions (resolved / locked) & alternatives

**D1 — New Rayls handler, not stock Circle bytecode (LOCKED).** Rationale: [§3](#3-why-not-integrate-circle-directly). Alternatives rejected: raw `deploy(bytecode,…)` of stock Circle runtime (factory still calls the Rayls `initialize` selector → revert); interface-detection in Blockscout instead of event-based (divergent detection path, more surface than emitting one event).

**D2 — Generic `deployRegistered` over a typed `deployStableCoin` wrapper (phase 1, recommended).** `PackDeployRegistered(key, userArgs, resourceId)` already exists in the generated binding (`RNContractFactoryV1.go:458`). Calling it with `key = keccak256("RAYLS_STABLECOIN")`, `userArgs = abi.encode(name, symbol, decimals)` deploys the type with **no binding regen, no factory redeploy**. The typed wrapper is nicer ABI but needs contract + interface edits, a factory redeploy, and binding regen. *Implication:* the generic registered path still emits `RegisteredContractDeployed`, which `extractDeployedAddress` already unpacks (`rayls_contract_factory.go:174`) — no event-handling change.

**D3 — `userArgs` = `(string name, string symbol, uint8 decimals)` (LOCKED).** Unchanged from ERC-20; masterMinter / pauser / blacklister all default to `trusted.owner`. **The `POST /api/tokens` body needs no new fields.** If a future phase needs distinct role addresses at deploy, that's a separate change extending the tuple + DTO.

**D4 — Phase-1 feature surface (LOCKED):** pause + blacklist + masterMinter/controllers (`configureMinter`/`removeMinter`, allowance-capped `mint`). **Deferred:** EIP-3009 (`transferWithAuthorization`), EIP-2612 (`permit`), rescuer.

**D5 — ops-api enum value `7`.** Next free after `ErcStandardZkDvpERC1155 = 6`. `erc_standard` is already `smallint` → no migration. The on-chain `RaylsBridgeableERC` tag from `GetERCStandard()` is separate — phase 1 may reuse the ERC20 tag for teleport metadata (the type distinction lives in the creation event + factory key), unless a new bridge tag is explicitly required.

---

## 7. Tasks

> Sequencing: §7.1–7.2 (contracts) are the blocking foundation. §7.3 (Blockscout) and §7.4 (ops-api) can be built in parallel but cannot be verified end-to-end until 7.1–7.2 are deployed/seeded on the target PN.

### 7.0 Scope (LOCKED)
- [x] **0.1** Feature subset: pause, blacklist, masterMinter + controlled `configureMinter`/`mint`/`burn`. EIP-3009/2612/rescuer **deferred**.
- [x] **0.2** `userArgs` = `(string name, string symbol, uint8 decimals)`; roles default to `trusted.owner`; `POST /api/tokens` body unchanged.

### 7.1 Contracts — handler & factory key (rayls-privacy-contracts)
- [ ] **1.1** Add `bytes32 internal constant RAYLS_STABLECOIN_KEY = keccak256("RAYLS_STABLECOIN");` to `FactoryKeys.sol`.
- [ ] **1.2** Expose `bytes32 public constant RAYLS_STABLECOIN_KEY = FactoryKeys.RAYLS_STABLECOIN_KEY;` on `AbstractContractFactoryV1`.
- [ ] **1.3** Create `RaylsStableCoinHandler.sol` extending `RaylsErc20Handler`. Surface: **pause** (`pause()`/`unpause()` gating transfers), **blacklist** (`blacklist(addr)`/`unBlacklist(addr)`), **masterMinter + controllers** (`configureMinter(minter, allowance)`/`removeMinter(minter)`, allowance-capped `mint`). Emit `event RaylsStableCoinCreated(address indexed tokenAddress);` from the overridden `initialize` — NOT `RaylsErc20TokenCreated`.
- [ ] **1.4** Override `initialize(bytes userArgs, RaylsTrustedInit trusted)`: decode `(name, symbol, decimals)`, call base init, set `masterMinter = pauser = blacklister = trusted.owner`, register new restricted selectors (`pause`/`unpause`/`blacklist`/`unBlacklist`/`configureMinter`/`removeMinter`) via the `_registerAccessControl` pattern, emit `RaylsStableCoinCreated`. *Decide:* keep the base owner-`mint` AND add controller-mint, or gate base mint behind the minter-allowance check.
- [ ] **1.5** Override `teleportDeployHint()` → `keccak256("RAYLS_STABLECOIN")` + matching `userArgs` (so receiver-side teleport materializes the stablecoin, not a plain ERC-20).
- [ ] **1.6** *(Optional, decision-gated)* Add typed `deployStableCoin(string,string,uint8,bytes32)` to `IBaseContractFactory.sol` + `AbstractContractFactoryV1.sol` → `_deployRegistered(RAYLS_STABLECOIN_KEY, abi.encode(name, symbol, decimals), resourceId)`. Skip if using generic `deployRegistered`.

### 7.2 Contracts — production contract & seeding (rayls-privacy-contracts)
- [ ] **2.1** Create `ProductionStableCoin.sol` as a thin concrete subclass of `RaylsStableCoinHandler` (mirroring `ProductionErc20Token.sol`).
- [ ] **2.2** Add `{ label: 'StableCoin', keyFn: 'RAYLS_STABLECOIN_KEY', artifact: 'ProductionStableCoin' }` to the standards table in `hardhat/tasks/deploy/privacy-node.ts` (the `setBytecode(KEY, deployedBytecode)` block).
- [ ] **2.3** Verify compiled runtime is under the `InitCodeStub` max length (else `InitCodeStub__RuntimeTooLarge`). Trim surface if needed.
- [ ] **2.4** Unit tests: register `RAYLS_STABLECOIN_KEY`; assert `deployRegistered(key, abi.encode(name, symbol, decimals), 0)` deploys, emits `RegisteredContractDeployed`, instance emits `RaylsStableCoinCreated`, correct name/symbol/decimals, pause/blacklist/mint behave.
- [ ] **2.5** *(If 1.6 taken)* Regenerate the `RNContractFactoryV1` abigen binding in ops-api so `PackDeployStableCoin` exists. Otherwise no binding change.

### 7.3 Blockscout detection (rayls-privacy-blockscout)
- [ ] **3.1** Add `{"RaylsStableCoinCreated(address)", "Rayls-StableCoin"}` to `@event_signatures` in `rayls_token_discovery.ex`.
- [ ] **3.2** Update the hardcoded `[t1..t6] = Map.keys(topic_to_type)` destructure and the `where: l.first_topic in [^t1, ...]` clause to include the 7th topic — or refactor to pin N topics. Each topic MUST be individually pinned (`in ^list` skips `Hash.Full` casting and silently matches nothing).
- [ ] **3.3** Confirm `initial_last_scanned_block` (`fragment("? LIKE 'Rayls-%'")`) already covers `'Rayls-StableCoin'` (prefix matches; verify).

### 7.4 ops-api mapping (rayls-sovereign-ops-api)
- [ ] **4.1** Add `ErcStandardStableCoin ErcStandard = 7` to the const block in `domain/types.go`; add `case ErcStandardStableCoin: return "RAYLS_STABLECOIN"` to `Label()`.
- [ ] **4.2** Add `case "Rayls-StableCoin": return ErcStandardStableCoin` (+ canonical-label `"RAYLS_STABLECOIN"`) to the shared `domain.ParseErcStandard`. *The prefix-mismatch fix already landed (see §5); this is one more `case`, no second call site.*
- [ ] **4.3** Add deploy case in `deployCalldata`: `case domain.ErcStandardStableCoin: return s.factory.PackDeployRegistered(stablecoinKey, abi.encode(name, symbol, decimals), resourceID), nil`, `stablecoinKey = crypto.Keccak256Hash([]byte("RAYLS_STABLECOIN"))`. (Use `PackDeployStableCoin` if 1.6 taken.)
- [ ] **4.4** Accept the standard in `parseErcStandard` + `supportedStandards()` in `token_deploy_handler.go`.
- [ ] **4.5** *(Only if 0.2 changes)* — N/A under locked scope; `ercStandardFromString` already routes through `domain.ParseErcStandard` (4.2).
- [ ] **4.6** Regenerate Swagger if the deploy DTO changed (it doesn't, under locked scope); confirm `ercStandardLabel = "RAYLS_STABLECOIN"` in responses.

### 7.5 End-to-end verification
- [ ] **5.1** On a seeded test PN, `POST /api/tokens` with the stablecoin standard; confirm deploy succeeds and the receipt has `RegisteredContractDeployed` under `RAYLS_STABLECOIN_KEY`.
- [ ] **5.2** Confirm Blockscout writes `tokens.type = 'Rayls-StableCoin'`.
- [ ] **5.3** Confirm `GET /api/tokens/:address` returns `ercStandardLabel = "RAYLS_STABLECOIN"` and the token appears in the SSE stream with the new standard.
- [ ] **5.4** Confirm mint/burn permissions resolve (AccessManager-derived callable functions) for the stablecoin owner.

---

## 8. Risks / trade-offs

- **Contracts work is the critical path.** ops-api + Blockscout edits are inert until the handler is written, the production contract compiled, the key seeded, and the factory holds the bytecode on each PN.
- **`InitCodeStub__RuntimeTooLarge`.** Pause/blacklist/minter logic grows the runtime; the factory wraps via `InitCodeStub.wrapRuntimeMemory` (max length). Keep the surface lean; check `deployedBytecode` size before seeding.
- **Blockscout 6→7 topic pin.** The fetcher destructures exactly six topics because `in ^list` skips casting. A silent miss here means the token deploys but is never classified.
- **No Circle audit lineage.** The re-implementation is your own contract; budget an audit for it (see §3 caveat).

---

## 9. Migration plan

1. Land contracts: handler + production contract + key + seeding; `setBytecode(RAYLS_STABLECOIN_KEY, runtime)` on target PNs.
2. Land Blockscout fetcher change; redeploy the indexer so new deploys classify as `Rayls-StableCoin`.
3. Land ops-api mapping; deploy. `POST /api/tokens` with the stablecoin standard now deploys via `deployRegistered` and is detected.
4. Verify end-to-end (§7.5).

---

## 10. Spec delta — `tokens` capability

### ADDED Requirement: Rayls stablecoin token standard
The system SHALL recognize a new ERC standard `RAYLS_STABLECOIN` (enum value `7`, `domain.ErcStandardStableCoin`) representing a Rayls-native, Circle-style stablecoin deployed through the RNContractFactory. It SHALL be deployable via `POST /api/tokens`, classified by the Blockscout indexer, and exposed in token list/detail/SSE responses with `ercStandardLabel = "RAYLS_STABLECOIN"`. `domain.ErcStandard.Label()` SHALL return `"RAYLS_STABLECOIN"` for value `7`. The stablecoin SHALL be deployed as a Rayls handler (`RaylsStableCoinHandler`) registered under `keccak256("RAYLS_STABLECOIN")`, NOT as stock Circle `FiatTokenV2` bytecode, and the instance SHALL emit `RaylsStableCoinCreated(address)` so the fetcher can classify it.

- **Scenario — Deploy via the tokens API:** WHEN an authorized caller issues `POST /api/tokens` with the stablecoin standard and `{name, symbol, decimals}` THEN the system SHALL build calldata targeting the `RAYLS_STABLECOIN` registry key (`deployRegistered(keccak256("RAYLS_STABLECOIN"), abi.encode(name, symbol, decimals), bytes32(0))`, or the typed wrapper), sign via the caller's custody wallet, submit to the RNContractFactory, and return the deployed address + tx hash on success.
- **Scenario — Deploy revert surfaces as an error:** WHEN the deploy tx reverts (e.g. the bytecode isn't seeded on the target PN) THEN the system SHALL return an error and MUST NOT report success or persist an internal token row as deployed.
- **Scenario — Stablecoin deploy is stored internal:** WHEN deployed through `POST /api/tokens` THEN it SHALL be stored `TokenStatusInternal` and SHALL NOT be promoted to `Active` when the indexer later discovers it.

### ADDED Requirement: Stablecoin classification from Blockscout
The integration SHALL classify a token whose Blockscout `tokens.type` is `"Rayls-StableCoin"` as `domain.ErcStandardStableCoin`. The `RaylsTokenDiscovery` fetcher SHALL map the `RaylsStableCoinCreated(address)` topic to `tokens.type = "Rayls-StableCoin"`.

- **Scenario — Indexer maps the type string:** WHEN the indexer processes a token whose Blockscout `type` is `"Rayls-StableCoin"` THEN `parseErcStandard` SHALL return `domain.ErcStandardStableCoin` and the served `erc_standard` SHALL be `7`.
- **Scenario — Unknown types remain custom:** WHEN the type matches no recognized standard THEN it SHALL classify as `ErcStandardCustom` (`"CUSTOM"`), unchanged by this addition.

### ADDED Requirement: Stablecoin standard in deploy and query surfaces
The deploy handler SHALL accept the stablecoin standard and list it among supported standards. The repository SHALL map the standard string in both directions for list/filter queries.

- **Scenario — Deploy handler accepts the standard:** WHEN a `POST /api/tokens` request specifies the stablecoin standard THEN `parseErcStandard` SHALL resolve it to `domain.ErcStandardStableCoin` and `supportedStandards()` SHALL include it; an unsupported string SHALL still be rejected.
- **Scenario — Filter list by standard:** WHEN a caller lists tokens filtered by the stablecoin standard THEN the repository SHALL resolve it to `erc_standard = 7` and return only stablecoin tokens.
