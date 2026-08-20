#!/usr/bin/env bash
#
# start_dev.sh — minimal dev stack for the ops-service + playground.
#
# Brings up a local dev stack for the ops-service + playground on top of a local RayUp
# control plane (api + worker on a k3d cluster, ../rayup). No chain is provisioned at
# startup — you click "Create Sovereign Chain" in the playground and RayUp spins one up
# on demand, with its own ops-api, contracts and explorer.
#
# The host stack (custody, nats, postgres, playground) comes up chain-less; login and
# admin bootstrap need no chain. Everything that needs a live chain is provided per chain
# by RayUp when you create one.
#
# Usage:
#   ./start_dev.sh [options]
#
# Options:
#   --rayup               Bring up the RayUp control plane (../rayup) on a local k3d cluster
#                         plus the ops stack, CHAIN-LESS. No VPN needed. No chain is
#                         provisioned — you create one from the playground ("Create Sovereign
#                         Chain") and RayUp spins it up on demand. First run builds the axyl
#                         image and creates the cluster (slow); later runs reuse both.
#   --clean               (--rayup) FULL RESET, then start fresh. Destroys every RayUp chain
#                         and its explorer, then tears down everything: the ops stack, the
#                         RayUp control plane, the shared Blockscout infra, the k3d cluster
#                         AND its registry, plus the images and volumes they built. Replaces
#                         the old reset-dev.sh (and "--rayup --down"). Everything is then
#                         rebuilt from scratch, including the contracts-deployer image
#                         (Go + Foundry + hardhat), so a chain you create afterwards gets
#                         its contracts with no second command.
#                         SLOW by design — this is the "nothing survives" spelling. A plain
#                         --rayup keeps the cluster, registry, images and caches so restarts
#                         stay fast; use that unless you actually want a clean slate.
#   --redeploy            Force a fresh contract deploy (ignore the cached registry).
#   --no-playground       Don't start the playground service.
#   --bootstrap-email <e> Admin email for POST /admin/bootstrap (overrides OPS_ADMIN_EMAIL).
#   -h, --help            Show this help.

set -euo pipefail

# macOS still ships bash 3.2 as /bin/bash, and under `set -u` that version treats an EMPTY
# array expansion (`"${ARR[@]}"`) as an unbound variable and aborts. This script uses empty
# arrays in several places — `--no-playground` leaves PROFILE_ARGS empty, for one — so on a
# stock Mac it would die partway through with a baffling "unbound variable".
#
# Fail immediately with something actionable instead. Homebrew's bash is 5.x and
# `#!/usr/bin/env bash` picks it up once it is on PATH.
if [ "${BASH_VERSINFO[0]:-0}" -lt 4 ]; then
    printf '\033[31mError:\033[0m bash >= 4 required (found %s at %s).\n' \
        "${BASH_VERSION:-unknown}" "${BASH:-bash}" >&2
    printf '  macOS ships bash 3.2. Install a newer one and re-run:\n' >&2
    printf '    brew install bash\n' >&2
    exit 1
fi

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

COMPOSE_FILE="docker-compose.dev-remote.yml"
REMOTE_ENV_DIR="docker/development"
REMOTE_ENV_FILE="$REMOTE_ENV_DIR/remote.env"
REMOTE_ENV_EXAMPLE="$REMOTE_ENV_DIR/remote.env.example"
# Public-chain mode caches its state in a separate file so a PN run and a public-test
# run never clobber each other's registry/key (the active file is chosen below).
PUBLIC_TEST_ENV_FILE="$REMOTE_ENV_DIR/public-test.env"
# RayUp mode caches its state (chain RPC/id, instance id, registry) in its own file too,
# so a --pn, a --deploy-public-test and a --rayup run never clobber each other.
RAYUP_ENV_FILE="$REMOTE_ENV_DIR/rayup.env"
CONTRACTS_DIR="../rayls-privacy-contracts"
PLAYGROUND_DIR="../rayls-privacy-playground"
CUSTODY_DIR="../rayls-privacy-custody-light"
BLOCKSCOUT_DIR="../rayls-privacy-blockscout"
RELAYER_DIR="../rayls-privacy-relayer-api"
RAYUP_DIR="../rayup"
DEPLOY_LOG="$REMOTE_ENV_DIR/contracts-deploy.log"

# Defaults
REDEPLOY=false
WITH_PLAYGROUND=true
DO_CLEAN=false
BOOTSTRAP_EMAIL_OVERRIDE=""
# This is a RayUp-only local dev stack (chain created on demand from the playground).
# RAYUP / CHAINLESS_RUN are constants, not selectable modes.
RAYUP=true
CHAINLESS_RUN=true

# --- RayUp mode ------------------------------------------------------------------
# RayUp's control plane (api + worker on a k3d cluster) provisions the CHAIN and
# nothing else — its own contracts/explorer/ops/playground phases stay off, because
# this script owns those (and runs them with hot reload, which k8s can't give us).
RAYUP_API_URL="${RAYUP_API_URL:-http://localhost:4000}"
# Headless auth: a script has no browser, so it can't do RayUp's OAuth round-trip.
# Must match DEV_API_TOKEN in RayUp's control-plane compose.
RAYUP_API_TOKEN="${RAYUP_API_TOKEN:-dev-token}"
# Service token for the playground's server-side rayup proxy. Unlike RAYUP_API_TOKEN
# (the shared-admin handshake token used by THIS script), the proxy sends this ALONG
# WITH the caller's ops-api email, so RayUp scopes chains PER USER — otherwise every
# playground account collapses to the shared dev admin and sees everyone's chains.
# The same value must reach both the control plane and the playground (below).
RAYUP_SERVICE_TOKEN="${RAYUP_SERVICE_TOKEN:-dev-service-token}"
# Emails RayUp treats as ADMIN (they see every chain, not just their own). Everyone
# else is scoped to the chains they created. Defaults to the bootstrap admin.
RAYUP_ADMIN_EMAILS="${RAYUP_ADMIN_EMAILS:-${OPS_ADMIN_EMAIL:-admin@example.com}}"
RAYUP_INSTANCE_NAME="${RAYUP_INSTANCE_NAME:-ops-dev}"
RAYUP_COMPOSE=(docker compose -f docker-compose.dev.yml --profile control-plane)
# The contracts-deployer image RayUp's worker runs (as a one-off in-cluster Job) right
# after a new chain is RUNNING. Built from the contracts repo (docker/rayup/Dockerfile).
#
# TWO NAMES, ONE REGISTRY — the same split dev-cluster.sh uses for the axyl image:
# we PUSH to localhost:5000 (reachable from this host) but REFERENCE it as
# k3d-<name>:5000, which is what the CLUSTER resolves via its containerd mirror.
RAYUP_REGISTRY_NAME="${REGISTRY_NAME:-rayls-registry}"
RAYUP_REGISTRY_PORT="${REGISTRY_PORT:-5000}"
RAYUP_DEPLOYER_REPO="rayls-contracts-deployer"
RAYUP_DEPLOYER_TAG="dev"
RAYUP_DEPLOYER_PUSH="localhost:${RAYUP_REGISTRY_PORT}/${RAYUP_DEPLOYER_REPO}:${RAYUP_DEPLOYER_TAG}"
RAYUP_DEPLOYER_IMAGE="k3d-${RAYUP_REGISTRY_NAME}:${RAYUP_REGISTRY_PORT}/${RAYUP_DEPLOYER_REPO}:${RAYUP_DEPLOYER_TAG}"

# The per-chain ops stack images (ops-api + ops-worker), built from THIS repo and run
# in-cluster by RayUp's ops-instance chart — one stack per chain, exactly as in cloud.
# Same push/reference split as the deployer above.
RAYUP_OPS_TAG="dev"
RAYUP_OPS_API_REPO="rayls-ops-api"
RAYUP_OPS_WORKER_REPO="rayls-ops-worker"
RAYUP_OPS_CUSTODY_REPO="rayls-custody"
RAYUP_OPS_API_PUSH="localhost:${RAYUP_REGISTRY_PORT}/${RAYUP_OPS_API_REPO}:${RAYUP_OPS_TAG}"
RAYUP_OPS_WORKER_PUSH="localhost:${RAYUP_REGISTRY_PORT}/${RAYUP_OPS_WORKER_REPO}:${RAYUP_OPS_TAG}"
RAYUP_OPS_CUSTODY_PUSH="localhost:${RAYUP_REGISTRY_PORT}/${RAYUP_OPS_CUSTODY_REPO}:${RAYUP_OPS_TAG}"
RAYUP_OPS_API_IMAGE="k3d-${RAYUP_REGISTRY_NAME}:${RAYUP_REGISTRY_PORT}/${RAYUP_OPS_API_REPO}:${RAYUP_OPS_TAG}"
RAYUP_OPS_WORKER_IMAGE="k3d-${RAYUP_REGISTRY_NAME}:${RAYUP_REGISTRY_PORT}/${RAYUP_OPS_WORKER_REPO}:${RAYUP_OPS_TAG}"
RAYUP_OPS_CUSTODY_IMAGE="k3d-${RAYUP_REGISTRY_NAME}:${RAYUP_REGISTRY_PORT}/${RAYUP_OPS_CUSTODY_REPO}:${RAYUP_OPS_TAG}"

# The per-chain Blockscout backend (RayUp's blockscout-instance chart runs one explorer per
# chain, in-cluster, alongside that chain's ops stack). Built from the Blockscout repo's
# production Dockerfile — the Rayls fork, so it must be OURS and not the upstream image.
# The frontend is stock ghcr.io/blockscout/frontend and is pulled straight from the chart.
RAYUP_BLOCKSCOUT_REPO="rayls-blockscout"
RAYUP_BLOCKSCOUT_PUSH="localhost:${RAYUP_REGISTRY_PORT}/${RAYUP_BLOCKSCOUT_REPO}:${RAYUP_OPS_TAG}"
RAYUP_BLOCKSCOUT_IMAGE="k3d-${RAYUP_REGISTRY_NAME}:${RAYUP_REGISTRY_PORT}/${RAYUP_BLOCKSCOUT_REPO}:${RAYUP_OPS_TAG}"

# Shared bases the per-chain ops stacks reuse in dev. They run as HOST containers already
# (the blockscout shared stack + this repo's compose), published on these host ports, so
# pods reach them through the k3d gateway instead of us standing up in-cluster copies.
# Each chain still gets its own ops_api_<slug>/raylzdb_<slug> databases on that server —
# per-chain DATA isolation is the point, not a per-chain Postgres.
# The host, as seen FROM INSIDE THE CLUSTER. k3d writes this name into CoreDNS's NodeHosts
# on every platform, so it is answered authoritatively by the cluster itself.
#
# NOT host.docker.internal: that name is not in NodeHosts — it only resolves because
# CoreDNS falls through to the node's resolver, which happens to be Docker Desktop's. That
# works here and on macOS, but it depends on the Docker distribution rather than on k3d,
# and it is not answered at all under a plain dockerd. host.k3d.internal has neither
# caveat, and resolves to the same address.
OPS_DEV_SHARED_HOST="host.k3d.internal"
OPS_DEV_POSTGRES_PORT="7432"      # shared-db (blockscout/docker-compose/shared.yml)
OPS_DEV_NATS_PORT="4222"
OPS_DEV_CUSTODY_PORT="5032"       # custody (docker-compose.dev-remote.yml "5032:5000")
OPS_DEV_REDIS_PORT="7379"         # shared-redis (blockscout/docker-compose/shared.yml)
# Phoenix signing key for the per-chain Blockscout backends. The same DEV value the
# host-side compose explorers use (blockscout/docker-compose/envs/common-blockscout.env),
# so both halves agree; it only signs local explorer sessions.
BLOCKSCOUT_DEV_SECRET_KEY_BASE="dev-only-not-a-secret-dev-only-not-a-secret-dev-only-not-a-secret"
# DEV-ONLY credentials, identical to the compose stack's so both halves agree.
SHARED_DB_PASSWORD="dev-only-not-a-secret"
OPS_DEV_JWT_SECRET="dev-secret-do-not-use-in-prod"
OPS_DEV_CUSTODY_API_KEY="dev-only-custody-api-key"
OPS_DEV_CUSTODY_JWT_SECRET="dev-only-not-a-secret-32byte-min"
# The KEYSTORE password custody encrypts wallet keys with. MUST equal the value the
# identity service mints with (CUSTODY_RAYLS_HSM_PASSWORD in docker-compose.dev-remote.yml):
# with one shared custody, a wallet minted under one password cannot be decrypted under
# another — signing fails with "Cannot derive the same mac as the one provided".
OPS_DEV_CUSTODY_PASSWORD="dev-password"
# Provisioning is genesis + N validators + an RPC health gate — minutes, not seconds.
RAYUP_CREATE_TIMEOUT="${RAYUP_CREATE_TIMEOUT:-900}"
# Destroy is `helm uninstall` + namespace teardown, which k8s does not do instantly.
RAYUP_DESTROY_TIMEOUT="${RAYUP_DESTROY_TIMEOUT:-90}"

log()  { printf '\033[36m==>\033[0m %s\n' "$*"; }
ok()   { printf '\033[32m ✓\033[0m %s\n' "$*"; }
warn() { printf '\033[33m ! \033[0m%s\n' "$*" >&2; }
die()  { printf '\033[31mError:\033[0m %s\n' "$*" >&2; exit 1; }

# The help text IS the header comment block: everything from line 2 up to the first
# line that isn't a comment (`set -euo pipefail`). Deriving the end rather than
# hardcoding it means adding an option can't silently truncate --help.
usage() { awk 'NR>1 { if (!/^#/) exit; sub(/^# ?/, ""); print }' "$0"; exit 0; }

# Big-integer helpers for wei amounts.
#
# Wei values exceed 2^53, so neither bash arithmetic (64-bit — it overflows to a NEGATIVE
# number) nor awk (doubles — a 1-wei difference rounds to 0) can be trusted here. `bc`
# would work but is not guaranteed on macOS, and a missing bc fails SILENTLY: the
# comparison yields "" and the shortfall becomes empty, so `cast send --value ""` dies with
# an unrelated error. python3 is exact, and ships with the Command Line Tools that cast and
# git already require.
wei_gte() { python3 -c "import sys; sys.exit(0 if int(sys.argv[1]) >= int(sys.argv[2]) else 1)" "$1" "$2"; }
wei_gt()  { python3 -c "import sys; sys.exit(0 if int(sys.argv[1]) >  int(sys.argv[2]) else 1)" "$1" "$2"; }
wei_sub() { python3 -c "import sys; print(int(sys.argv[1]) - int(sys.argv[2]))" "$1" "$2"; }

# ------------------------------------------------------------------ arg parsing
while [ $# -gt 0 ]; do
    case "$1" in
        --rayup)         shift ;;   # accepted for compatibility; rayup is the only mode
        --clean)         DO_CLEAN=true; shift ;;
        # Refused rather than silently ignored: this script no longer provisions a chain
        # (the playground does, on demand), so a developer expecting "this makes me a fresh
        # chain" gets corrected instead of quietly getting the cached one.
        --new-chain)
            die "--new-chain is gone: chains are created from the playground now (\"Create Sovereign
     Chain\"), not by this script. Use --rayup --clean for a fresh slate, then create a chain in the UI." ;;
        --redeploy)      REDEPLOY=true; shift ;;
        --no-playground) WITH_PLAYGROUND=false; shift ;;
        --bootstrap-email) BOOTSTRAP_EMAIL_OVERRIDE="${2:?--bootstrap-email requires a value}"; shift 2 ;;
        -h|--help)       usage ;;
        *) die "unknown option: $1 (use --help)" ;;
    esac
done

# The RayUp chain is provisioned locally and its RPC/chainId are only known AFTER the
# instance is RUNNING (see provision_rayup_chain below, which writes them into rayup.env).
# The host stack boots chain-less until then; everything that needs a live chain is gated
# on CHAINLESS_RUN and skipped here. --clean is the teardown spelling (it destroys the
# RayUp chains and the whole stack).
REMOTE_ENV_FILE="$RAYUP_ENV_FILE"
log "Mode: RAYUP — local axyl chain on k3d (created on demand from the playground)"

# --------------------------------------------------------------------- env file
# Read a KEY=VALUE from an env file (no export, ignores comments).
# Always returns 0 — a missing key yields an empty string (so `VAR="$(read_env …)"`
# under `set -e -o pipefail` doesn't abort when the key is absent).
read_env() {
    local key="$1" file="$2"
    [ -f "$file" ] || return 0
    grep -E "^${key}=" "$file" 2>/dev/null | tail -n1 | cut -d= -f2- || true
}

# Idempotently set KEY=VALUE in an env file (append or replace in place).
set_env_var() {
    local file="$1" key="$2" value="$3" tmp
    tmp="$(mktemp)"
    if [ -f "$file" ] && grep -qE "^${key}=" "$file"; then
        grep -vE "^${key}=" "$file" > "$tmp"
        printf '%s=%s\n' "$key" "$value" >> "$tmp"
        cat "$tmp" > "$file"
    else
        printf '%s=%s\n' "$key" "$value" >> "$file"
    fi
    rm -f "$tmp"
}

# ------------------------------------------------------------------ RayUp (chain)
# RayUp is the control plane that provisions axyl privacy chains onto a Kubernetes
# cluster. In --rayup mode we point it at a LOCAL k3d cluster and ask it for one
# chain; everything above the chain stays here. RayUp's own contracts/explorer/ops/
# playground phases are disabled in its compose, so it stops once the chain is up.

# Call the RayUp API. Prints the response body on stdout; returns non-zero on a
# network failure OR an HTTP 4xx/5xx.
#
# Plain `curl -sS` exits 0 on 4xx/5xx, so a 500 would flow on as a "response" and only
# surface later as a baffling parse error ("RayUp did not return an instance id:
# {"error":...}"). We want the status in the exit code but the BODY on stdout — that
# is exactly --fail-with-body (curl >= 7.76). Bare -f would give us the exit code and
# throw the body away, which loses the one thing that explains the failure.
#
# On older curl (--fail-with-body unsupported) fall back to asking for the status code
# separately and failing on it ourselves, so the behaviour is the same everywhere.
if curl --help all 2>/dev/null | grep -q -- '--fail-with-body'; then
    CURL_FAIL=(--fail-with-body)
else
    CURL_FAIL=()
fi

rayup_api() {
    local method="$1" path="$2"; shift 2
    local body rc arg
    local ctype=()

    # Send Content-Type ONLY when we are actually sending a body. RayUp's Fastify rejects
    # a bodyless request that declares application/json ("Body cannot be empty when
    # content-type is set to 'application/json'", HTTP 400) — which silently broke every
    # GET and, worse, every DELETE (chains were never destroyed, they just 400'd).
    for arg in "$@"; do
        case "$arg" in
            -d|--data|--data-raw|--data-binary) ctype=(-H 'Content-Type: application/json'); break ;;
        esac
    done

    if [ ${#CURL_FAIL[@]} -gt 0 ]; then
        curl -sS "${CURL_FAIL[@]}" -X "$method" "${RAYUP_API_URL}${path}" \
            -H "Authorization: Bearer ${RAYUP_API_TOKEN}" \
            "${ctype[@]}" "$@"
        return $?
    fi

    # Fallback: append the status code, split it off, fail on >= 400.
    body="$(curl -sS -w '\n%{http_code}' -X "$method" "${RAYUP_API_URL}${path}" \
        -H "Authorization: Bearer ${RAYUP_API_TOKEN}" \
        "${ctype[@]}" "$@")" || return $?
    rc="${body##*$'\n'}"
    printf '%s' "${body%$'\n'*}"
    [ "$rc" -lt 400 ] 2>/dev/null || return 22   # 22 == curl's HTTP-error exit code
}

# Resolve the deploy key into $DEPLOY_KEY, creating one if this is a first run.
#
# PN/rayup mode auto-generates a per-developer throwaway key (those chains are gasless).
# Public mode uses a stable seeded key, topped up from the faucet before deploy. Callable
# more than once — rayup needs the key BEFORE the control plane starts (the deployer Job
# signs with it), and the deploy step needs it later.
ensure_deploy_key() {
    [ -n "${DEPLOY_KEY:-}" ] && return 0

    DEPLOY_KEY="$(read_env PRIVATE_KEY_SYSTEM "$REMOTE_ENV_FILE")"
    [ -n "$DEPLOY_KEY" ] && return 0

    log "Generating a per-developer deploy key (cast wallet new)…"
    local new_wallet deploy_addr
    new_wallet="$(cast wallet new)"
    DEPLOY_KEY="$(echo "$new_wallet" | awk '/Private key:/ {print $3}')"
    deploy_addr="$(echo "$new_wallet" | awk '/Address:/ {print $2}')"
    [ -n "$DEPLOY_KEY" ] || die "could not parse 'cast wallet new' output"
    set_env_var "$REMOTE_ENV_FILE" "PRIVATE_KEY_SYSTEM" "$DEPLOY_KEY"
    ok "Deploy key created (address $deploy_addr)."
}

# Is the contracts-deployer image in the cluster registry? Asked against the registry's
# own catalog rather than `docker images`: what matters is whether the CLUSTER can pull
# it, and a local daemon copy proves nothing about that.
rayup_deployer_image_present() {
    curl -sf --max-time 3 "http://localhost:${RAYUP_REGISTRY_PORT}/v2/${RAYUP_DEPLOYER_REPO}/tags/list" 2>/dev/null \
        | grep -q "\"${RAYUP_DEPLOYER_TAG}\""
}

# Build the contracts-deployer image and push it to the cluster registry.
#
# Two layers, both from the contracts repo: the production image (which already bakes the
# compiled artifacts, TypeChain typings and the hub-less deploy tasks), then a thin layer
# swapping the entrypoint for RayUp's deploy-rayup.sh — the script that runs the hub-less
# privacy-node deploy and prints the ---RAYUP_RESULT--- payload the worker parses.
build_rayup_deployer_image() {
    [ -d "$CONTRACTS_DIR" ] || die "missing sibling repo: $CONTRACTS_DIR (needed to build the deployer)"
    [ -f "$CONTRACTS_DIR/docker/rayup/Dockerfile" ] \
        || die "no $CONTRACTS_DIR/docker/rayup/Dockerfile — this contracts checkout has no RayUp deployer."

    log "Building the contracts image (Go + Foundry + hardhat — several minutes on a cold cache)…"
    ( cd "$CONTRACTS_DIR" && docker build -t rayls-contracts:local . ) \
        || die "failed to build rayls-contracts:local (see the build output above)"

    log "Building the RayUp deployer layer…"
    ( cd "$CONTRACTS_DIR" && docker build -f docker/rayup/Dockerfile -t "$RAYUP_DEPLOYER_PUSH" . ) \
        || die "failed to build $RAYUP_DEPLOYER_PUSH"

    log "Pushing to the cluster registry…"
    docker push "$RAYUP_DEPLOYER_PUSH" || die "failed to push $RAYUP_DEPLOYER_PUSH — is the k3d registry up? (k3d registry list)"

    # Tag under the in-cluster name too, mirroring what dev-cluster.sh does for axyl.
    docker tag "$RAYUP_DEPLOYER_PUSH" "$RAYUP_DEPLOYER_IMAGE" 2>/dev/null || true
    ok "Deployer image available in-cluster as $RAYUP_DEPLOYER_IMAGE"
}

# Create the cluster secrets every per-chain ops stack reads (rayup-root namespace).
#
# The ops-instance chart expects two Secrets that, in cloud, are applied by hand once
# (k8s/system/*.example.yaml). Locally we generate them from the SAME dev values the
# compose stack already uses, so the per-chain ops-apis share the host's Postgres and
# custody credentials rather than needing a second set.
#
# Idempotent: recreated on every run so an edited .env.oauth takes effect.
ensure_rayup_ops_secrets() {
    command -v kubectl >/dev/null 2>&1 || { warn "kubectl not found — skipping ops secrets."; return 0; }

    kubectl create namespace rayup-root >/dev/null 2>&1 || true

    # Postgres password for the shared DB (published on the host by the blockscout
    # shared stack); the chart reads it as POSTGRES_PASSWORD from this secret.
    #
    # SECRET_KEY_BASE is the Blockscout backend's Phoenix signing key. The per-chain
    # explorer chart mounts it from THIS secret, so without it the backend never starts
    # ("couldn't find key SECRET_KEY_BASE") and the token/balance lists stay empty — the
    # ops-api mirrors that explorer's database. Fixed dev value: it only signs local
    # sessions, and a stable one keeps them valid across restarts.
    kubectl create secret generic blockscout-shared \
        --namespace rayup-root \
        --from-literal=POSTGRES_PASSWORD="$SHARED_DB_PASSWORD" \
        --from-literal=SECRET_KEY_BASE="$BLOCKSCOUT_DEV_SECRET_KEY_BASE" \
        --dry-run=client -o yaml 2>/dev/null | kubectl apply -f - >/dev/null 2>&1 \
        || { warn "could not create the blockscout-shared secret."; return 0; }

    # ops-api session key + custody credentials + the shared Google OAuth app. The
    # custody values MUST match what the ops-api presents (same pair as compose).
    kubectl create secret generic ops-shared \
        --namespace rayup-root \
        --from-literal=OPS_JWT_SECRET="$OPS_DEV_JWT_SECRET" \
        --from-literal=CUSTODY_API_KEY="$OPS_DEV_CUSTODY_API_KEY" \
        --from-literal=CUSTODY_JWT_SECRET="$OPS_DEV_CUSTODY_JWT_SECRET" \
        --from-literal=CUSTODY_PASSWORD="$OPS_DEV_CUSTODY_PASSWORD" \
        --from-literal=GOOGLE_CLIENT_ID="${GOOGLE_CLIENT_ID:-}" \
        --from-literal=GOOGLE_CLIENT_SECRET="${GOOGLE_CLIENT_SECRET:-}" \
        --dry-run=client -o yaml 2>/dev/null | kubectl apply -f - >/dev/null 2>&1 \
        || { warn "could not create the ops-shared secret."; return 0; }

    ok "Cluster secrets ready (blockscout-shared, ops-shared)."
}

# Are ALL THREE per-chain ops images already in the cluster registry?
rayup_ops_images_present() {
    local repo
    for repo in "$RAYUP_OPS_API_REPO" "$RAYUP_OPS_WORKER_REPO" "$RAYUP_OPS_CUSTODY_REPO"; do
        curl -sf --max-time 3 "http://localhost:${RAYUP_REGISTRY_PORT}/v2/${repo}/tags/list" 2>/dev/null \
            | grep -q "\"${RAYUP_OPS_TAG}\"" || return 1
    done
    return 0
}

# Is the per-chain Blockscout backend image in the cluster registry? Checked separately
# from the ops images: a missing explorer costs the token/balance lists (they mirror its
# database), but the chain and its ops-api are still usable, so it must not gate them.
rayup_blockscout_image_present() {
    curl -sf --max-time 3 "http://localhost:${RAYUP_REGISTRY_PORT}/v2/${RAYUP_BLOCKSCOUT_REPO}/tags/list" 2>/dev/null \
        | grep -q "\"${RAYUP_OPS_TAG}\""
}

# Build the per-chain ops-api + ops-worker images and push them to the cluster registry.
#
# These are what RayUp's ops-instance chart runs: ONE ops stack per chain, each bound to
# its own chain and its own database. That per-chain isolation is the whole point — a
# chain's tokens/balances/roles live in ITS ops-api, and both die with the chain, so a new
# chain can never inherit a previous one's state.
#
# Built from this repo's production Dockerfiles (same binaries as cloud, `run` and
# `worker` entrypoints, migrations baked in so each per-chain DB self-migrates on boot).
# Push one image to the cluster registry, retrying a few times.
#
# The registry is local and unauthenticated, but `docker push` still consults the
# configured credential helper. Under Docker Desktop on WSL that helper is a Windows
# binary reached over vsock, which intermittently times out:
#
#   WSL ERROR: UtilAcceptVsock:273: accept4 failed 110
#   error getting credentials - err: exit status 1
#
# The helper answers again moments later, so the push is retried rather than failing the
# whole run — it is a transport flake between WSL and Docker Desktop, not a registry fault
# (the giveaway is earlier images in the SAME loop pushing fine).
push_to_cluster_registry() {
    local ref="$1" attempt
    for attempt in 1 2 3; do
        docker push "$ref" && return 0
        [ "$attempt" -lt 3 ] && {
            warn "push of $ref failed (attempt $attempt/3) — retrying in 5s…"
            sleep 5
        }
    done
    return 1
}

build_rayup_ops_images() {
    [ -d "$CUSTODY_DIR" ] || die "missing sibling repo: $CUSTODY_DIR (needed for the per-chain custody)"

    log "Building the ops-api image…"
    docker build -f Dockerfile -t "$RAYUP_OPS_API_PUSH" . \
        || die "failed to build $RAYUP_OPS_API_PUSH (see the build output above)"

    log "Building the ops-worker image…"
    docker build -f Dockerfile.worker -t "$RAYUP_OPS_WORKER_PUSH" . \
        || die "failed to build $RAYUP_OPS_WORKER_PUSH (see the build output above)"

    # Custody is per-chain too: it signs and broadcasts to ONE chain (fixed EIP-155 chain
    # id + RPC), so it cannot be shared between chains.
    log "Building the custody image…"
    ( cd "$CUSTODY_DIR" && docker build -f Dockerfile -t "$RAYUP_OPS_CUSTODY_PUSH" . ) \
        || die "failed to build $RAYUP_OPS_CUSTODY_PUSH (see the build output above)"

    log "Pushing the ops images to the cluster registry…"
    local push
    for push in "$RAYUP_OPS_API_PUSH" "$RAYUP_OPS_WORKER_PUSH" "$RAYUP_OPS_CUSTODY_PUSH"; do
        push_to_cluster_registry "$push" \
            || die "failed to push $push after 3 attempts — is the k3d registry up? (k3d registry list)"
    done

    # Tag under the in-cluster names too (what the chart's image refs use).
    docker tag "$RAYUP_OPS_API_PUSH" "$RAYUP_OPS_API_IMAGE" 2>/dev/null || true
    docker tag "$RAYUP_OPS_WORKER_PUSH" "$RAYUP_OPS_WORKER_IMAGE" 2>/dev/null || true
    docker tag "$RAYUP_OPS_CUSTODY_PUSH" "$RAYUP_OPS_CUSTODY_IMAGE" 2>/dev/null || true
    ok "Ops images available in-cluster (ops-api, ops-worker, custody)."
}

# Build + push the per-chain Blockscout backend (the Rayls fork) to the cluster registry.
#
# Separate from build_rayup_ops_images because it is a different repo and a much slower
# build (Elixir release), and because its absence is degraded-but-usable rather than fatal:
# the chain and ops-api still work, the token/balance lists just stay empty.
build_rayup_blockscout_image() {
    [ -d "$BLOCKSCOUT_DIR" ] || { warn "missing sibling repo: $BLOCKSCOUT_DIR — skipping the explorer image."; return 1; }

    log "Building the Blockscout backend image (Elixir release — this is slow on a cold cache)…"
    # Must go through the repo's build script, NOT a bare `docker build`: the Dockerfile
    # copies config_helper.exs to /app/releases/${RELEASE_VERSION}/, and runtime.exs loads
    # it from its own directory. Without the build-arg the file lands one level up and the
    # backend CrashLoopBackOffs at boot with "could not load .../config_helper.exs", which
    # shows up as an empty explorer AND empty token supply/holder counts (the ops-api
    # mirrors those from the explorer's DB). The script reads the version from mix.exs.
    ( cd "$BLOCKSCOUT_DIR" && REGISTRY="localhost:${RAYUP_REGISTRY_PORT}" TAG="$RAYUP_OPS_TAG" PUSH=false \
        ./scripts/build-image.sh ) \
        || { warn "failed to build $RAYUP_BLOCKSCOUT_PUSH — chains will be created without an explorer."; return 1; }

    log "Pushing the Blockscout image to the cluster registry…"
    push_to_cluster_registry "$RAYUP_BLOCKSCOUT_PUSH" \
        || { warn "failed to push $RAYUP_BLOCKSCOUT_PUSH after 3 attempts — is the k3d registry up? (k3d registry list)"; return 1; }

    docker tag "$RAYUP_BLOCKSCOUT_PUSH" "$RAYUP_BLOCKSCOUT_IMAGE" 2>/dev/null || true
    ok "Blockscout image available in-cluster (per-chain explorer)."
}

# Tear down every per-chain Blockscout instance we created (the rayup-<chainId> ones).
#
# Each is a 4-container stack plus a database in the shared-db, created on demand when the
# ops stack binds to a chain. Their chains are destroyed by --clean, which leaves them
# indexing an RPC that no longer answers and holding three host ports each — so they must
# go too, or they pile up across runs.
#
# -v drops the instance's volumes; the DB lives in the SHARED shared-db and survives that,
# so it is dropped explicitly. Both are worthless once the chain is gone.
destroy_rayup_blockscout_instances() {
    local dir name dropped=0

    for dir in "$BLOCKSCOUT_DIR"/docker-compose/instances/rayup-*/; do
        [ -d "$dir" ] || continue                      # no matches → the glob itself
        name="$(basename "$dir")"
        case "$name" in
            rayup) continue ;;                          # the hand-made one, not ours
        esac

        log "Removing Blockscout instance $name…"
        ( cd "$dir" && docker compose down -v --remove-orphans >/dev/null 2>&1 ) \
            || warn "could not stop Blockscout instance $name."

        if docker ps --format '{{.Names}}' | grep -qx shared-db; then
            docker exec shared-db psql -U blockscout -d postgres \
                -c "DROP DATABASE IF EXISTS \"blockscout_${name}\" WITH (FORCE);" >/dev/null 2>&1 \
                || warn "could not drop blockscout_${name}."
        fi

        rm -rf "$dir"
        dropped=$((dropped + 1))
    done

    if [ "$dropped" -gt 0 ]; then
        # The cached conn string may name a database we just dropped. Leaving it would hand
        # the worker a dead DSN and it would retry "database does not exist" forever, so the
        # pointer dies with the thing it points at.
        case "$(read_env BLOCKSCOUT_DB_CONN "$REMOTE_ENV_FILE")" in
            *blockscout_rayup-*) set_env_var "$REMOTE_ENV_FILE" "BLOCKSCOUT_DB_CONN" "" ;;
        esac
        ok "Removed $dropped per-chain Blockscout instance(s)."
    fi
    return 0
}

# Force-remove every container matching a `docker ps` query (args are passed straight to
# docker). No-op when nothing matches.
#
# Deliberately not `… | xargs -r docker rm -f`: -r ("skip if input is empty") is a GNU
# extension that BSD/macOS xargs rejects outright, so that spelling fails on a Mac AND
# runs `docker rm -f` with no arguments. Reading into a variable is portable.
docker_rm_matching() {
    local ids
    ids="$(docker "$@" 2>/dev/null)" || return 0
    [ -n "$ids" ] || return 0
    # shellcheck disable=SC2086  # intentional split: ids is a newline-separated list
    docker rm -f $ids >/dev/null 2>&1 || true
}

# Force-remove every image whose repo:tag matches an ERE. No-op when nothing matches.
docker_rmi_matching() {
    local pattern="$1" images
    images="$(docker images --format '{{.Repository}}:{{.Tag}}' 2>/dev/null | grep -E "$pattern")" || return 0
    [ -n "$images" ] || return 0
    # shellcheck disable=SC2086  # intentional split: images is a newline-separated list
    docker rmi -f $images >/dev/null 2>&1 || true
}

# Hard teardown of every Rayls dev stack: containers, networks, images and volumes.
#
# This is the second half of --clean (the first half, destroy_all_rayup_chains +
# destroy_rayup_blockscout_instances, must run BEFORE this while the control plane is
# still up — destroying a chain is a job its worker runs, and the per-chain Blockscout
# databases live in shared-db). By the time we get here those are already reaped, so what
# remains is the infrastructure itself.
#
# Absorbed from the old reset-dev.sh, which this replaces. Scoped by name to the rayls /
# rayup / blockscout / k3d stacks — unrelated Docker on the machine is left alone.
#
# A plain --rayup run never calls this: the k3d cluster, its registry (holding the built
# contracts-deployer + axyl images) and the compose images are exactly what make a restart
# fast. --clean is the "fresh slate" spelling, so here they all go and the normal startup
# path below rebuilds them.
teardown_everything() {
    local down_flags=(down --remove-orphans -v)
    local rayup_compose="$RAYUP_DIR/docker-compose.dev.yml"
    local shared_compose="$BLOCKSCOUT_DIR/docker-compose/shared.yml"
    local net vol

    log "Downing the ops stack…"
    # Both profiles: a down must reach every service regardless of which mode started it,
    # or profile-gated containers (ops-api/worker) survive the teardown as orphans.
    compose --profile playground --profile bound-chain "${down_flags[@]}" >/dev/null 2>&1 \
        || warn "could not fully down the ops stack — continuing."

    log "Downing the RayUp control plane…"
    [ -f "$rayup_compose" ] && { docker compose -f "$rayup_compose" --profile control-plane \
        "${down_flags[@]}" >/dev/null 2>&1 || warn "could not fully down the RayUp control plane."; }

    log "Downing the shared Blockscout infra…"
    [ -f "$shared_compose" ] && { docker compose -f "$shared_compose" "${down_flags[@]}" \
        >/dev/null 2>&1 || warn "could not fully down the shared Blockscout infra."; }

    # The cluster AND its registry. dev-cluster.sh --delete does both; the registry is a
    # separate k3d object that a bare `k3d cluster delete` leaves running (this is exactly
    # what reset-dev.sh missed, stranding k3d-rayls-registry after a "full" reset).
    log "Deleting the k3d cluster + registry…"
    if [ -x "$RAYUP_DIR/hack/dev-cluster.sh" ]; then
        ( cd "$RAYUP_DIR" && ./hack/dev-cluster.sh --delete ) >/dev/null 2>&1 \
            || warn "could not delete the k3d cluster — continuing."
    fi
    # Belt and braces: whatever the CLI left behind (or all of it, if k3d is not installed).
    docker_rm_matching ps -aq --filter 'name=k3d-rayls'

    # Stragglers a compose `down` misses — containers started outside their project
    # attribution (the per-chain rayup-<id>-* explorer stacks are the usual culprit).
    docker_rm_matching ps -aq --filter 'name=rayls-privacy-ops-api' --filter 'name=rayup' \
        --filter 'name=shared-'

    for net in blockscout-shared rayup_default rayls-privacy-ops-api_default \
               docker-compose_default k3d-rayls-dev; do
        docker network rm "$net" >/dev/null 2>&1 || true
    done

    # Images built by these compose projects, plus RayUp's per-chain Blockscout backends
    # (rayup-<chainId>-*). Stable names, so a grep is safe.
    log "Removing built images…"
    docker_rmi_matching '^(rayls-privacy-ops-api|rayup)([-_].*)?:'
    docker_rmi_matching '^rayup-[0-9]+-'

    # Named volumes the compose downs above could not reach (the playground's runtime state
    # among them — it holds the chain selection, which must not survive a clean).
    for vol in $(docker volume ls --format '{{.Name}}' 2>/dev/null \
                 | grep -E '^(rayls-privacy-ops-api|rayup)[-_]'); do
        docker volume rm -f "$vol" >/dev/null 2>&1 || true
    done
    docker volume prune -f >/dev/null 2>&1 || true

    ok "Teardown complete — containers: $(docker ps -aq 2>/dev/null | wc -l)."
    return 0
}

# NOTE: the host-side ops rebind (rebind_ops_stack / autoselect_new_chain /
# watch_rayup_selection) is gone. Each chain now gets its OWN ops-api, worker and
# Blockscout, provisioned with it by RayUp (helm releases ops-<slug> / bs-<slug>, each
# with its own database) and destroyed with it. Nothing needs rebinding, because nothing
# is shared: the playground routes each request to the ops-api of the chain it names
# (src/lib/chain-target.ts), ownership-checked against the caller.
#
# The old model bound ONE shared ops-api to ONE chain at a time. The second user to create
# a chain never got the slot, so their chain was running and deployed but unusable, and
# their requests were served by the first user's stack.

# Bring RayUp's control plane up: k3d cluster (+ axyl image) → container kubeconfig
# → postgres/redis/migrate/api/worker. Idempotent; reuses whatever already exists.
# Bring RayUp's control plane up. Pass "no-deployer-build" to skip the contracts-deployer
# image build: --clean calls this twice (once before the teardown, purely so its worker can
# destroy the existing chains, and again after to rebuild), and building on the first pass
# would burn several minutes on an image the teardown then deletes.
start_rayup_control_plane() {
    local skip_deployer_build=false
    [ "${1:-}" = "no-deployer-build" ] && skip_deployer_build=true

    log "Ensuring the k3d cluster + axyl image are ready (first run is slow)…"
    ( cd "$RAYUP_DIR" && ./hack/dev-cluster.sh ) || die "failed to prepare the k3d cluster ($RAYUP_DIR/hack/dev-cluster.sh)"

    # The worker runs in a container, so it needs a kubeconfig whose server address
    # resolves from inside one (k3d writes 0.0.0.0/127.0.0.1, which would mean the
    # worker container itself).
    ( cd "$RAYUP_DIR" && ./hack/kubeconfig-for-container.sh >/dev/null ) \
        || die "failed to write the container kubeconfig ($RAYUP_DIR/hack/kubeconfig-for-container.sh)"

    # --clean is a "fresh, fully-working slate", so it GUARANTEES the deployer image is
    # present — without it, created chains get no contracts and their dashboard stays
    # unavailable. --clean now deletes the registry along with the cluster, so this build
    # always runs on the post-teardown pass (several minutes, unavoidable: it IS the fresh
    # slate). Still gated on absence, for the no-op case where the registry survived. Must
    # come after dev-cluster.sh (the registry to push to exists only once the cluster is up)
    # and before the presence check just below that turns the feature on. A plain --rayup
    # run leaves an absent image absent (see the warn below) and reuses a present one.
    if [ "$DO_CLEAN" = true ] && [ "$skip_deployer_build" != true ] && ! rayup_deployer_image_present; then
        log "--clean: contracts-deployer image missing — building it so new chains get contracts…"
        build_rayup_deployer_image
    fi

    # The per-chain ops images are needed on EVERY rayup run, not just --clean: without
    # them a created chain gets no ops stack at all (RayUp logs "Ops stack skipped" and the
    # chain has no API behind its dashboard). Unlike the deployer these are OUR code and
    # change constantly, so a plain --rayup rebuilds them too — the Go layer cache makes a
    # no-change rebuild quick. Skipped on the pre-teardown pass of --clean, whose images the
    # teardown deletes moments later.
    if [ "$skip_deployer_build" != true ]; then
        log "Building the per-chain ops images (ops-api + ops-worker)…"
        build_rayup_ops_images

        # The explorer backend is NOT rebuilt every run: it is a slow Elixir release and it
        # is not our code changing day to day (unlike ops-api/worker). Build it only when the
        # registry does not already have it — i.e. first run, or after a --clean wiped it.
        if rayup_blockscout_image_present; then
            ok "Blockscout image already in the cluster registry."
        else
            build_rayup_blockscout_image || true
        fi
    fi

    # Contract deploy is part of CHAIN CREATION: RayUp's worker runs the deployer as a
    # one-off in-cluster Job right after a chain comes up, so a chain created from the
    # playground is usable immediately — no second command. It only does this when told
    # which image to run, which is what we pass here.
    #
    # Gated on the image actually being present in the cluster registry: without it every
    # create would enqueue a Job that cannot pull, and the chain would look broken. Absent,
    # we simply leave contracts off (RayUp logs "deployer not configured") and say so.
    if rayup_deployer_image_present; then
        ensure_deploy_key   # the Job signs the deploy with it — needed before we start
        export CONTRACTS_DEPLOYER_IMAGE="$RAYUP_DEPLOYER_IMAGE"
        export PRIVATE_KEY_SYSTEM="$DEPLOY_KEY"
        # Wired into the contracts as a value, never dialed (DEPLOY_PUBLIC_CHAIN=false).
        export PUBLIC_CHAIN_ID="${PUBLIC_CHAIN_ID:-1337}"
        export DEPLOY_PUBLIC_CHAIN="false"
        ok "Contract deploy enabled for new chains ($RAYUP_DEPLOYER_IMAGE)."
    else
        export CONTRACTS_DEPLOYER_IMAGE=""
        warn "Contracts deployer image not found in the cluster registry — chains will be
     created WITHOUT contracts (the chain dashboard stays unavailable). Build it with:
       ./start_dev.sh --rayup --clean"
    fi

    # Per-chain ops stack. Each chain gets its OWN ops-api + worker + custody, created
    # with the chain and destroyed with it — the same topology as cloud, and the reason
    # a new chain can no longer inherit a previous chain's tokens or registry.
    #
    # Gated on the images being in the registry for the same reason as the deployer: a
    # missing image would leave every chain with an un-pullable ops stack.
    if rayup_ops_images_present; then
        export OPS_ENABLED="true"
        # No ingress/cert-manager on k3d — publish each ops-api on a host NodePort.
        export OPS_DEV_MODE="true"
        export OPS_DEV_EXTERNAL_HOST="localhost"
        # How the RayUp WORKER (a container) reaches those same host ports.
        export OPS_DEV_WORKER_HOST="host.docker.internal"
        export OPS_DEV_PLAYGROUND_ORIGIN="http://localhost:3000"
        # Reuse the host's Postgres/NATS (published ports) rather than in-cluster copies.
        export OPS_DEV_POSTGRES_HOST="$OPS_DEV_SHARED_HOST"
        export OPS_DEV_POSTGRES_PORT="$OPS_DEV_POSTGRES_PORT"
        export OPS_DEV_NATS_URL="nats://${OPS_DEV_SHARED_HOST}:${OPS_DEV_NATS_PORT}"
        # Every chain signs through the SAME custody the identity service mints with: a user
        # has ONE custody wallet across all chains, and only the HSM that minted it holds the
        # key. A per-chain custody would 204 on every signing request ("wallet not found").
        export OPS_DEV_CUSTODY_URL="http://${OPS_DEV_SHARED_HOST}:${OPS_DEV_CUSTODY_PORT}"
        # The chart's secrets must exist before any chain is created.
        ensure_rayup_ops_secrets
        # RayUp bootstraps each new chain's ops-api with an admin. The instance OWNER is
        # used first (the logged-in user who created the chain); this is the fallback.
        # Resolved here rather than read from the global OPS_ADMIN_EMAIL, which is not
        # computed until much later in the script — long after the control plane starts.
        if [ -z "${OPS_ADMIN_EMAIL:-}" ]; then
            for _oauth in ".env.oauth" "$RELAYER_DIR/.env.oauth"; do
                [ -f "$_oauth" ] || continue
                OPS_ADMIN_EMAIL="$(read_env OPS_ADMIN_EMAIL "$_oauth")"
                [ -n "$OPS_ADMIN_EMAIL" ] && break
            done
        fi
        export OPS_ADMIN_EMAIL="${BOOTSTRAP_EMAIL_OVERRIDE:-${OPS_ADMIN_EMAIL:-}}"
        ok "Per-chain ops stack enabled (each chain gets its own ops-api + DB)."

        # Per-chain explorer, on the same terms as the ops stack: created with the chain,
        # destroyed with it. The ops-api's token/balance lists are mirrors of ITS database,
        # so a chain with no explorer shows no tokens — hence enabling it alongside ops.
        #
        # Gated on the image being in the registry for the same reason as ops: a missing
        # image would leave every chain with an un-pullable explorer.
        if rayup_blockscout_image_present; then
            export EXPLORER_ENABLED="true"
            # As with ops: no ingress/cert-manager on k3d, so publish ONE host NodePort per
            # explorer (the chart fronts backend+frontend with its own nginx, doing the path
            # split the Traefik ingress does in cloud).
            export EXPLORER_DEV_MODE="true"
            # A third disjoint slice of the k3d host-mapped band: chain RPC counts UP from
            # 30000, ops DOWN from 30100, explorers DOWN from 30050.
            export EXPLORER_DEV_NODEPORT_MAX="30050"
            # Reuse the host's Redis (published by the shared Blockscout compose) — the
            # cluster has no blockscout-redis of its own. Postgres comes from the ops vars.
            export EXPLORER_DEV_REDIS_HOST="$OPS_DEV_SHARED_HOST"
            export EXPLORER_DEV_REDIS_PORT="$OPS_DEV_REDIS_PORT"
            ok "Per-chain explorer enabled (each chain gets its own Blockscout + DB)."
        else
            export EXPLORER_ENABLED="false"
            warn "Blockscout image not in the cluster registry — chains will be created
     WITHOUT an explorer (token and balance lists will stay empty)."
        fi
    else
        export OPS_ENABLED="false"
        # The explorer only earns its keep alongside an ops stack (its database is what the
        # ops-api's token/balance lists mirror), so it follows ops off.
        export EXPLORER_ENABLED="false"
        warn "Ops images not found in the cluster registry — chains will be created
     WITHOUT an ops stack."
    fi

    log "Starting the RayUp control plane (postgres, redis, api, worker)…"
    ( cd "$RAYUP_DIR" && \
        DEV_API_TOKEN="$RAYUP_API_TOKEN" \
        RAYUP_SERVICE_TOKEN="$RAYUP_SERVICE_TOKEN" \
        ADMIN_EMAILS="$RAYUP_ADMIN_EMAILS" \
        OPS_ENABLED="$OPS_ENABLED" \
        OPS_DEV_MODE="${OPS_DEV_MODE:-false}" \
        OPS_DEV_EXTERNAL_HOST="${OPS_DEV_EXTERNAL_HOST:-localhost}" \
        OPS_DEV_WORKER_HOST="${OPS_DEV_WORKER_HOST:-host.docker.internal}" \
        OPS_DEV_PLAYGROUND_ORIGIN="${OPS_DEV_PLAYGROUND_ORIGIN:-http://localhost:3000}" \
        OPS_OAUTH_REDIRECT_BASE="${OPS_OAUTH_REDIRECT_BASE:-http://localhost:3000}" \
        OPS_DEV_POSTGRES_HOST="${OPS_DEV_POSTGRES_HOST:-}" \
        OPS_DEV_POSTGRES_PORT="${OPS_DEV_POSTGRES_PORT:-}" \
        OPS_DEV_NATS_URL="${OPS_DEV_NATS_URL:-}" \
        OPS_DEV_CUSTODY_URL="${OPS_DEV_CUSTODY_URL:-}" \
        OPS_ADMIN_EMAIL="${OPS_ADMIN_EMAIL:-}" \
        EXPLORER_ENABLED="${EXPLORER_ENABLED:-false}" \
        EXPLORER_DEV_MODE="${EXPLORER_DEV_MODE:-false}" \
        EXPLORER_DEV_NODEPORT_MAX="${EXPLORER_DEV_NODEPORT_MAX:-30050}" \
        EXPLORER_DEV_REDIS_HOST="${EXPLORER_DEV_REDIS_HOST:-}" \
        EXPLORER_DEV_REDIS_PORT="${EXPLORER_DEV_REDIS_PORT:-}" \
        "${RAYUP_COMPOSE[@]}" up -d --build ) \
        || die "failed to start the RayUp control plane"

    log "Waiting for the RayUp API ($RAYUP_API_URL/health)…"
    local elapsed=0
    while [ $elapsed -lt 180 ]; do
        curl -sf -o /dev/null --max-time 2 "$RAYUP_API_URL/health" && { ok "RayUp API up."; return 0; }
        sleep 3; elapsed=$((elapsed + 3))
    done
    die "RayUp API did not become healthy at $RAYUP_API_URL — check: (cd $RAYUP_DIR && ${RAYUP_COMPOSE[*]} logs api)"
}

# Destroy EVERY chain the control plane knows about, waiting for each to actually land.
#
# Chains are created from the playground UI, not by this script, so there is no cached
# instance id to target — we enumerate. Used by --clean (wipe before starting fresh); a
# DELETE only ENQUEUES the job, so we wait for each (404/DELETED) before returning, or the
# next step could kill the worker mid-`helm uninstall` and leak a namespace + NodePort.
destroy_all_rayup_chains() {
    local ids iid destroy_err

    if ! curl -sf -o /dev/null --max-time 3 "$RAYUP_API_URL/health"; then
        warn "RayUp API not reachable at $RAYUP_API_URL — skipping chain teardown. Any existing
     chain may leak its namespace + NodePort."
        return 0
    fi

    ids="$(rayup_api GET /api/instances 2>/dev/null \
        | jq -r '.instances[]? | select(.status != "DELETED") | .id' 2>/dev/null || true)"
    if [ -z "$ids" ]; then
        ok "No RayUp chains to destroy."
        return 0
    fi

    for iid in $ids; do
        log "Destroying RayUp chain $iid (frees its namespace + NodePort)…"
        if destroy_err="$(rayup_api DELETE "/api/instances/$iid" 2>&1)"; then
            wait_rayup_destroyed "$iid" || true
        else
            warn "could not destroy chain $iid (${destroy_err:-no response}) — do it
     manually or it will leak its NodePort."
        fi
    done
}

# Poll an instance to RUNNING, echoing each phase transition ("Running genesis",
# "Starting validators", …). Dies with the job log on ERROR or timeout.
wait_rayup_instance() {
    local id="$1" elapsed=0 body status phase last_phase="" bad=0

    while [ "$elapsed" -lt "$RAYUP_CREATE_TIMEOUT" ]; do
        # A transient failure here (API restarting mid-poll) is not fatal — but a
        # PERSISTENT one means we are polling a dead API, and that must not look the
        # same as "still provisioning". Provisioning takes minutes, so a developer
        # watching a frozen prompt has no way to tell the two apart unless we say so.
        if body="$(rayup_api GET "/api/instances/$id")"; then
            bad=0
        else
            bad=$((bad + 1))
            # ~30s of consecutive failures: say something, keep trying.
            [ $((bad % 6)) -eq 0 ] \
                && warn "RayUp API not answering for /api/instances/$id (${bad} tries): ${body:-no response}"
            sleep 5; elapsed=$((elapsed + 5))
            continue
        fi

        status="$(echo "$body" | jq -r '.instance.status // empty' 2>/dev/null || true)"
        phase="$(echo "$body" | jq -r '.instance.phase // empty' 2>/dev/null || true)"

        # 200 but no status = a body we do not understand (schema drift, a proxy's
        # error page). Silently retrying on that would burn the full timeout.
        if [ -z "$status" ]; then
            bad=$((bad + 1))
            [ $((bad % 6)) -eq 0 ] \
                && warn "RayUp returned no instance status (${bad} tries): $body"
        else
            bad=0
        fi

        if [ -n "$phase" ] && [ "$phase" != "$last_phase" ]; then
            printf '     %s…\n' "$phase"
            last_phase="$phase"
        fi

        case "$status" in
            RUNNING) ok "Chain is RUNNING."; return 0 ;;
            ERROR)
                warn "RayUp reported ERROR while provisioning. Job log:"
                rayup_api GET "/api/instances/$id/logs" | jq -r '.logs // empty' >&2 || true
                die "chain provisioning failed"
                ;;
        esac
        sleep 5; elapsed=$((elapsed + 5))
    done
    die "chain did not reach RUNNING within ${RAYUP_CREATE_TIMEOUT}s (see: ${RAYUP_COMPOSE[*]} logs worker in $RAYUP_DIR)"
}

# Poll a destroy to completion. RayUp flips the instance to DESTROYING, the worker runs
# `helm uninstall` + drops the namespace, and the row is soft-deleted — at which point
# GET returns 404. So "gone" is provable (404 or DELETED); we do not have to guess how
# long helm needs.
wait_rayup_destroyed() {
    local id="$1" elapsed=0 body status rc bad=0 slug

    while [ "$elapsed" -lt "$RAYUP_DESTROY_TIMEOUT" ]; do
        body="$(rayup_api GET "/api/instances/$id")" && rc=0 || rc=$?

        # An HTTP error (rayup_api normalizes those to curl's 22) is the SUCCESS signal
        # we are waiting for: the instance is soft-deleted, so the API 404s on it.
        # Any OTHER non-zero rc is curl failing to *reach* the API (7 refused, 28 timeout,
        # …) — that proves nothing about the chain, and must not be read as "destroyed"
        # just because the API happened to restart mid-teardown.
        if [ "$rc" -eq 22 ]; then
            ok "Chain $id destroyed."
            return 0
        elif [ "$rc" -ne 0 ]; then
            bad=$((bad + 1))
            # ~30s of consecutive transport failures: say something, keep polling.
            [ $((bad % 10)) -eq 0 ] \
                && warn "RayUp API not reachable while awaiting destroy of $id (${bad} tries, curl rc=$rc)"
            sleep 3; elapsed=$((elapsed + 3))
            continue
        fi
        bad=0

        status="$(echo "$body" | jq -r '.instance.status // empty' 2>/dev/null || true)"
        if [ "$status" = "DELETED" ]; then
            ok "Chain $id destroyed."
            return 0
        fi
        sleep 3; elapsed=$((elapsed + 3))
    done

    slug="$(read_env RAYUP_SLUG "$RAYUP_ENV_FILE")"
    warn "chain $id still ${status:-unknown} after ${RAYUP_DESTROY_TIMEOUT}s — it may leak its
     namespace + NodePort. Raise RAYUP_DESTROY_TIMEOUT, or clean up with:
       kubectl delete ns rayup-${slug:-<slug from RAYUP_SLUG in $RAYUP_ENV_FILE>}"
    return 1
}

# Is the chain RayUp calls RUNNING actually on the cluster? Answers only from the cluster
# itself (RayUp's DB survives a k3d wipe and keeps claiming RUNNING; see the caller).
#
# Deliberately CONSERVATIVE: returns "live" (0) unless it can positively prove the
# workloads are gone. A missing kubectl or an unreachable cluster is not proof — it must
# never destroy a healthy chain over a diagnostic hiccup. Only an existing, reachable
# cluster with no pods in the chain's namespace counts as dead.
rayup_chain_is_live() {
    local slug="$1" ns

    [ -n "$slug" ] || return 0                                  # can't tell → assume live
    command -v kubectl >/dev/null 2>&1 || return 0              # can't tell → assume live
    kubectl cluster-info >/dev/null 2>&1 || return 0            # can't tell → assume live

    ns="rayup-${slug}"                                          # RayUp names it rayup-<slug>
    kubectl get ns "$ns" >/dev/null 2>&1 || return 1            # namespace gone → dead
    [ -n "$(kubectl get pods -n "$ns" --no-headers 2>/dev/null)" ] || return 1  # no pods → dead
    return 0
}

# RayUp reported this chain as RUNNING but we cannot reach its RPC. Work out WHY and
# die with the actual cause. Ordered most-likely first: RayUp's DB outlives the k3d
# cluster, so recreating the cluster leaves a RUNNING row pointing at a namespace that
# no longer exists — which looks nothing like a NodePort problem but used to be
# reported as one.
rayup_unreachable_reason() {
    local ns="$1" url="$2"

    if ! command -v kubectl >/dev/null 2>&1; then
        die "cannot reach $url — the RayUp chain is RUNNING but unreachable, and kubectl
    is not installed so we cannot diagnose it. Check: kubectl get pods -n $ns"
    fi

    # The cluster itself must answer before any pod/namespace check means anything —
    # otherwise "no unhealthy pods" is indistinguishable from "no cluster".
    if ! kubectl cluster-info >/dev/null 2>&1; then
        die "cannot reach $url — the k3d cluster is not reachable (kubectl cannot talk to
    the API server), yet RayUp's database still says this chain is RUNNING. The cluster
    was deleted out from under it. Recreate the cluster, then start again and create a
    fresh chain from the playground (\"Create Sovereign Chain\"):
        (cd $RAYUP_DIR && ./hack/dev-cluster.sh)
        ./start_dev.sh --rayup --clean"
    fi

    [ -n "$ns" ] || die "cannot reach $url — and could not determine the chain's namespace,
    so we cannot diagnose it. Inspect it manually: kubectl get pods -A | grep rayup"

    if ! kubectl get ns "$ns" >/dev/null 2>&1; then
        die "cannot reach $url — RayUp says this chain is RUNNING, but its namespace
    ($ns) does not exist. The k3d cluster was recreated/deleted after the chain was
    provisioned, so RayUp's database is now stale. Start from a clean slate, then create
    a fresh chain from the playground (\"Create Sovereign Chain\"):
        ./start_dev.sh --rayup --clean"
    fi

    local pods not_running
    pods="$(kubectl get pods -n "$ns" --no-headers 2>/dev/null || true)"
    if [ -z "$pods" ]; then
        die "cannot reach $url — the namespace ($ns) exists but has no pods. The chain's
    workloads are gone while RayUp still reports it RUNNING. Start from a clean slate, then
    create a fresh chain from the playground (\"Create Sovereign Chain\"):
        ./start_dev.sh --rayup --clean"
    fi

    not_running="$(echo "$pods" | grep -viE 'Running|Completed' || true)"
    if [ -n "$not_running" ]; then
        die "cannot reach $url — the chain's pods are not healthy:
$not_running
    An ImagePullBackOff here means the axyl image is not pullable by its IN-CLUSTER
    name. Check the active ImageVersion is k3d-<registry>:5000/axyl:dev (not
    localhost:5000, which only works for pushing from the host):
        kubectl -n $ns describe pod <pod>"
    fi

    die "cannot reach $url — the chain's pods are healthy, so its NodePort is likely not
    published on the host. hack/dev-cluster.sh maps only 30000-\$NODEPORT_MAP_MAX;
    check that port ${RAYUP_PORT:-?} falls in that range."
}

# Create (or reuse) the chain and write its RPC/chainId into rayup.env. Sets the
# globals the rest of the script consumes.
provision_rayup_chain() {
    local id body status chain_id_hex chain_id_dec rpc_url slug port stale_chain_id cached_slug

    id="$(read_env RAYUP_INSTANCE_ID "$REMOTE_ENV_FILE")"

    # Reuse a cached instance only if it is genuinely still RUNNING (the cluster may
    # have been wiped, or the instance destroyed, since we last wrote the cache).
    #
    # RayUp's status alone is NOT enough: its Postgres lives on a docker volume that
    # survives a k3d wipe, so after the cluster is recreated the row still says RUNNING
    # while the namespace it points at is long gone. Trusting it yields a chain whose RPC
    # nothing answers. So verify against the CLUSTER — the only real source of truth —
    # and self-heal by provisioning a fresh chain when the workloads aren't there.
    if [ -n "$id" ]; then
        body="$(rayup_api GET "/api/instances/$id" 2>/dev/null || true)"
        status="$(echo "$body" | jq -r '.instance.status // empty' 2>/dev/null || true)"
        cached_slug="$(echo "$body" | jq -r '.instance.slug // empty' 2>/dev/null || true)"
        if [ "$status" != "RUNNING" ]; then
            warn "Cached RayUp chain $id is ${status:-gone} — provisioning a new one."
            id=""
        elif ! rayup_chain_is_live "$cached_slug"; then
            warn "Cached RayUp chain $id says RUNNING but its workloads are gone from the
     cluster (RayUp's DB outlived the k3d cluster) — provisioning a new one."
            id=""
        else
            ok "Reusing RayUp chain $id (RUNNING). Use --clean for a fresh slate."
        fi
    fi

    if [ -z "$id" ]; then
        # A NEW chain starts with empty state, so every pointer cached from the previous
        # one is dead: its DeploymentProxyRegistry address holds no code (ops-api refuses
        # to boot: "no contract code at address …") and its OZ manifest is keyed to a
        # chain that no longer exists. Drop them so the contracts get deployed fresh.
        stale_chain_id="$(read_env PRIVACY_NODE_CHAIN_ID "$REMOTE_ENV_FILE")"
        if [ -n "$stale_chain_id" ]; then
            rm -f "$CONTRACTS_DIR/.openzeppelin/unknown-${stale_chain_id}.json"
        fi
        set_env_var "$REMOTE_ENV_FILE" "DEPLOYMENT_PROXY_REGISTRY_ADDR" ""
        set_env_var "$REMOTE_ENV_FILE" "BLOCKCHAIN_STARTING_BLOCK" ""
        set_env_var "$REMOTE_ENV_FILE" "BUSINESS_ROLES_ACTIVATED" ""
        REGISTRY_ADDR=""

        log "Creating a RayUp chain (\"$RAYUP_INSTANCE_NAME\")…"
        body="$(rayup_api POST /api/instances -d "{\"name\":\"${RAYUP_INSTANCE_NAME}\"}")" \
            || die "POST /api/instances failed"
        id="$(echo "$body" | jq -r '.instance.id // empty')"
        [ -n "$id" ] || die "RayUp did not return an instance id: $body"
        ok "Instance $id queued — provisioning (genesis + validators takes a few minutes)…"
        wait_rayup_instance "$id"
    fi

    # Read back the provisioned chain.
    body="$(rayup_api GET "/api/instances/$id")"
    chain_id_hex="$(echo "$body" | jq -r '.instance.chainId // empty')"
    rpc_url="$(echo "$body" | jq -r '.instance.rpcUrl // empty')"
    slug="$(echo "$body" | jq -r '.instance.slug // empty')"
    port="$(echo "$body" | jq -r '.instance.rpcPort // empty')"
    # Kept for the reachability diagnostic below (see rayup_unreachable_reason).
    # The API does not serialize .instance.namespace (it is null on the wire), so derive
    # it from the slug — RayUp names the namespace (and the Helm release) "rayup-<slug>".
    RAYUP_NAMESPACE="$(echo "$body" | jq -r '.instance.namespace // empty')"
    [ -n "$RAYUP_NAMESPACE" ] || RAYUP_NAMESPACE="rayup-${slug}"
    RAYUP_PORT="$port"
    [ -n "$chain_id_hex" ] && [ -n "$port" ] || die "RayUp instance $id has no chainId/rpcPort: $body"

    # RayUp stores the chain id as hex; everything downstream (contracts, ops-api,
    # Blockscout, the playground) wants decimal.
    chain_id_dec="$((chain_id_hex))"

    # RayUp's own rpcUrl is built from the WORKER's NODE_PUBLIC_HOST (host.docker.internal
    # — the worker polls the RPC from inside its container). We need both forms: the host
    # one for our curl/cast/hardhat calls, the container one for the compose services.
    PN_RPC_URL="http://localhost:${port}"
    PN_RPC_URL_CONTAINER="http://host.docker.internal:${port}"

    set_env_var "$REMOTE_ENV_FILE" "RAYUP_INSTANCE_ID"      "$id"
    set_env_var "$REMOTE_ENV_FILE" "RAYUP_SLUG"             "$slug"
    set_env_var "$REMOTE_ENV_FILE" "PRIVACY_NODE_RPC_URL"   "$PN_RPC_URL"
    set_env_var "$REMOTE_ENV_FILE" "PRIVACY_NODE_CHAIN_ID"  "$chain_id_dec"
    # Feeless axyl chain, but the PN deploy task still wires a public chain id into the
    # RN endpoint init, so it must be set to something.
    [ -n "$(read_env PUBLIC_CHAIN_ID "$REMOTE_ENV_FILE")" ] \
        || set_env_var "$REMOTE_ENV_FILE" "PUBLIC_CHAIN_ID" "1337"

    ok "RayUp chain ready: $PN_RPC_URL (chainId $chain_id_dec, slug $slug)"
    echo "     RayUp instance: $id   (dashboard-less; manage via $RAYUP_API_URL/api/instances)"
}

# -------------------------------------------------- Blockscout instance selection
# Each mode gets its own Blockscout instance + port band so their explorers never
# collide. Needed by BOTH the teardown path (to stop the instance) and the startup
# path (to create/start it), so it is resolved here, before either.
#
# Each instance also claims <port+10> (stats) and <port+11> (visualizer), so the bases
# are spaced 100 apart. Bases: rayup 8440, public-test 8460, pn-a..pn-f 8480 + 100*index.
BS_INSTANCE="rayup"
BS_HTTP_PORT=8440
BS_INSTANCE_DIR="$BLOCKSCOUT_DIR/docker-compose/instances/$BS_INSTANCE"

# --------------------------------------------------------------------- tear down
compose() { docker compose -f "$COMPOSE_FILE" "$@"; }

# ------------------------------------------------------------------- preflight
log "Preflight checks…"
command -v docker >/dev/null 2>&1 || die "docker not found"
docker compose version >/dev/null 2>&1 || die "'docker compose' (v2) not found"
command -v curl >/dev/null 2>&1 || die "curl not found"
command -v cast >/dev/null 2>&1 || die "foundry 'cast' not found (needed to generate the deploy key) — https://book.getfoundry.sh/getting-started/installation"
command -v npx  >/dev/null 2>&1 || die "node/npx not found (needed for the contracts deploy)"
command -v jq   >/dev/null 2>&1 || warn "jq not found — admin bootstrap will be skipped (install: apt install jq / brew install jq)"

[ -d "$CONTRACTS_DIR" ]  || die "missing sibling repo: $CONTRACTS_DIR"
[ -d "$CUSTODY_DIR" ]    || die "missing sibling repo: $CUSTODY_DIR"
[ -d "$BLOCKSCOUT_DIR" ] || die "missing sibling repo: $BLOCKSCOUT_DIR (shared Blockscout)"
if [ "$WITH_PLAYGROUND" = true ]; then
    [ -d "$PLAYGROUND_DIR" ] || die "missing sibling repo: $PLAYGROUND_DIR (or pass --no-playground)"
fi
if [ "$RAYUP" = true ]; then
    command -v k3d >/dev/null 2>&1 || die "k3d not found (needed by --rayup) — https://k3d.io/#installation"
    # Unlike the other modes, --rayup parses the RayUp API's JSON, so jq is REQUIRED
    # (elsewhere it only gates the optional admin bootstrap).
    command -v jq >/dev/null 2>&1 || die "jq not found — required by --rayup (install: apt install jq / brew install jq)"
    [ -d "$RAYUP_DIR" ] || die "missing sibling repo: $RAYUP_DIR (needed by --rayup)"
fi
ok "Tooling and sibling repos present."

# ------------------------------------------------------ remote.env (create/load)
mkdir -p "$REMOTE_ENV_DIR"
if [ ! -f "$REMOTE_ENV_FILE" ]; then
    # Seed the cache from the example. The chain's RPC/chainId are unknown until it has
    # been provisioned, so they are filled in by provision_rayup_chain below.
    log "Creating $REMOTE_ENV_FILE from example (RayUp local chain)…"
    cp "$REMOTE_ENV_EXAMPLE" "$REMOTE_ENV_FILE"
fi

# ------------------------------------------------------------- RayUp: control plane
# --rayup brings up the RayUp CONTROL PLANE (k3d + api/worker on :4000) and provisions
# NO chain. The chain is created on demand when the user clicks "Create Sovereign Chain"
# in the playground (a POST /api/instances through the same-origin RayUp proxy).
#
# The rest of the ops stack (ops-api, worker, custody, nats, playground) STILL comes up —
# chain-less. It has to: login (Google/Microsoft/SIWE) and admin bootstrap are served by
# ops-api and need no chain (Bootstrap creates a custody wallet + DB rows only). What we
# skip below in this mode is everything that genuinely needs a live chain — reachability,
# contract deploy, business-roles, Blockscout — each gated on `$RAYUP` at its site.
if [ "$RAYUP" = true ]; then
    # The control plane decides whether to enable contract deploy by checking the registry
    # for the deployer image. --clean is a "fresh, fully-working slate", so it guarantees
    # that image is present (built if missing) before the check — see start_rayup_control_plane.
    #
    # Under --clean this first call exists only so the worker can destroy the existing
    # chains; everything it brings up is torn down moments later, so the deployer build is
    # deferred to the rebuild pass after the teardown.
    if [ "$DO_CLEAN" = true ]; then
        start_rayup_control_plane no-deployer-build
    else
        start_rayup_control_plane
    fi

    # --clean: wipe EVERYTHING, then fall through and start fresh.
    #
    # Two phases, and the order matters. The graceful phase needs the control plane alive
    # (destroying a chain is a job RayUp's worker runs, and the per-chain Blockscout DBs
    # live inside shared-db) — so chains and their explorers are reaped first, releasing
    # their host ports and dropping their databases. Only then does the hard phase remove
    # the infrastructure itself, control plane and k3d cluster included.
    #
    # Skipping straight to the hard teardown would orphan RayUp's port allocations and
    # leave blockscout_rayup-* databases behind in a shared-db that gets recreated empty.
    if [ "$DO_CLEAN" = true ]; then
        log "--clean: destroying every RayUp chain and tearing the whole stack down…"
        destroy_all_rayup_chains
        # LEGACY cleanup. Explorers are per-chain IN THE CLUSTER now (RayUp's bs-<slug>
        # release, destroyed with the chain), so nothing creates these host-side instances
        # any more. Kept so a machine that ran the older flow still gets its leftover
        # stacks, ports and databases reaped instead of leaking them forever. A no-op once
        # the directory is empty.
        destroy_rayup_blockscout_instances
        teardown_everything

        # The cluster we just deleted is the one the control plane talks to, so it has to
        # come back before anything else uses it. Rebuilds the k3d cluster + axyl image and
        # restarts postgres/redis/api/worker — the slow part of a --clean, and the price of
        # a genuinely fresh slate.
        log "Rebuilding the RayUp control plane from scratch…"
        start_rayup_control_plane

        ok "Clean slate. Starting fresh…"
    fi

    # --rayup boots chain-less: clear any chain pointers cached by a PREVIOUS rayup run.
    # That chain is gone (or was never ours), and stale values are actively harmful
    # chain-less — ops-api would take the old RPC/registry, try to verify the chain id
    # against a dead NodePort, and refuse to boot. Nothing repopulates them: in rayup mode
    # this stack is permanently chain-less, and each chain's own ops-api carries its config.
    set_env_var "$REMOTE_ENV_FILE" "PRIVACY_NODE_RPC_URL" ""
    set_env_var "$REMOTE_ENV_FILE" "PRIVACY_NODE_CHAIN_ID" ""
    set_env_var "$REMOTE_ENV_FILE" "DEPLOYMENT_PROXY_REGISTRY_ADDR" ""
    set_env_var "$REMOTE_ENV_FILE" "BLOCKCHAIN_STARTING_BLOCK" ""
    set_env_var "$REMOTE_ENV_FILE" "BUSINESS_ROLES_ACTIVATED" ""
    set_env_var "$REMOTE_ENV_FILE" "RAYUP_INSTANCE_ID" ""
    # And the per-chain explorer DB a previous rebind pointed us at: that database is
    # dropped along with its chain, so keeping the conn string would leave the worker
    # retrying a database that no longer exists, forever ("does not exist", backoff).
    # Empty is the correct chain-less value — it disables the Blockscout indexers.
    set_env_var "$REMOTE_ENV_FILE" "BLOCKSCOUT_DB_CONN" ""
fi

ensure_deploy_key

PN_RPC_URL="$(read_env PRIVACY_NODE_RPC_URL "$REMOTE_ENV_FILE")"
PN_CHAIN_ID="$(read_env PRIVACY_NODE_CHAIN_ID "$REMOTE_ENV_FILE")"
# The URL the CONTAINERS use to reach the chain. Same as PN_RPC_URL for a remote PN
# (a real public hostname); in rayup mode both stay empty — this stack is never bound to a
# chain, and each RayUp-provisioned chain's own ops-api holds its RPC URL.
PN_RPC_URL_CONTAINER="${PN_RPC_URL_CONTAINER:-$PN_RPC_URL}"
PUBLIC_CHAIN_ID="$(read_env PUBLIC_CHAIN_ID "$REMOTE_ENV_FILE")"
REGISTRY_ADDR="$(read_env DEPLOYMENT_PROXY_REGISTRY_ADDR "$REMOTE_ENV_FILE")"
STARTING_BLOCK="$(read_env BLOCKCHAIN_STARTING_BLOCK "$REMOTE_ENV_FILE")"
WALLETCONNECT_PROJECT_ID="$(read_env WALLETCONNECT_PROJECT_ID "$REMOTE_ENV_FILE")"
BLOCKSCOUT_DB_CONN="$(read_env BLOCKSCOUT_DB_CONN "$REMOTE_ENV_FILE")"

# OAuth (Google/Microsoft) creds for the playground login — same .env.oauth convention
# as the relayer. Checked in the repo root first, then the relayer's. Sourced with set -a
# so GOOGLE_CLIENT_ID/SECRET, MICROSOFT_CLIENT_ID/SECRET (and optional OPS_ADMIN_EMAIL)
# are exported to docker compose.
for _oauth in ".env.oauth" "$RELAYER_DIR/.env.oauth"; do
    if [ -f "$_oauth" ]; then
        set -a; . "$_oauth"; set +a
        ok "Loaded OAuth creds from $_oauth"
        break
    fi
done
OAUTH_ADMIN_EMAIL="${OPS_ADMIN_EMAIL:-}"   # OPS_ADMIN_EMAIL possibly set by .env.oauth
[ -n "${GOOGLE_CLIENT_ID:-}${MICROSOFT_CLIENT_ID:-}" ] || warn "No OAuth creds (.env.oauth) — Google/Microsoft login disabled (SIWE still works)."

# Admin email for bootstrap: flag wins, then remote.env, then .env.oauth.
OPS_ADMIN_EMAIL="${BOOTSTRAP_EMAIL_OVERRIDE:-$(read_env OPS_ADMIN_EMAIL "$REMOTE_ENV_FILE")}"
[ -n "$OPS_ADMIN_EMAIL" ] || OPS_ADMIN_EMAIL="$OAUTH_ADMIN_EMAIL"

[ -n "$PUBLIC_CHAIN_ID" ] || PUBLIC_CHAIN_ID=1337

# --rayup boots chain-less: PN_RPC_URL / PN_CHAIN_ID are legitimately empty (no chain
# until the UI creates one), so the empty-value guards and the reachability check that
# every OTHER mode requires must NOT run — they would die on the absent chain. Everything
# from here to the Blockscout section that needs a live chain is likewise gated on RAYUP.
log "RayUp mode: no chain until you create one in the playground (\"Create Sovereign Chain\")."

# Chain-less boot: the chain is a NodePort on THIS host that RayUp creates on demand, and
# each chain carries its own ops-api. No extra_hosts pin here.
PN_EXTRA_HOST=""

# ------------------------------------------------------------- contract deploy
# There is no chain to deploy to at boot: contracts are deployed by RayUp's own
# provisioning when the UI creates a chain, not by this script. (ops-api runs chain-less
# until then — its blockchain features self-disable when the registry is unset.)
log "RayUp mode: skipping contract deploy — a chain (and its contracts) is created from the UI."

# ------------------------------------------------------- activate business roles
# Business roles are (re)activated on a chain's AccessManager. In this RayUp stack that
# happens per chain when RayUp provisions it, not here.
log "RayUp mode: skipping business-roles activation — no contracts until a chain is created."

# ----------------------------------------------------- shared Blockscout (local)
# ops_api/raylzdb + the Blockscout DB all live in the shared-db of the
# `blockscout-shared` network. Bring the shared infra up and ensure a Blockscout
# instance is indexing this PN, so the ops worker can read tokens/balances.
# (BS_INSTANCE / BS_HTTP_PORT / BS_INSTANCE_DIR were resolved above the teardown block.)

# The shared infra (shared-db, redis) is NOT optional in any mode: ops_api and raylzdb —
# the databases ops-api, the worker and custody all connect to — live in that same
# shared-db, not just the Blockscout DBs. So this always runs, even chain-less.
log "Ensuring shared Blockscout infra is up…"
( cd "$BLOCKSCOUT_DIR" && ./scripts/start-shared.sh ) || die "failed to start shared Blockscout infra ($BLOCKSCOUT_DIR/scripts/start-shared.sh)"

# The per-chain Blockscout INSTANCE, by contrast, is pinned to ONE chain (its .env carries
# RAYLS_CHAIN_ID and a chain-guard service aborts the stack if the RPC disagrees), so it
# cannot exist until a chain does. In --rayup we boot chain-less, so skip it; the ops
# worker's token/balance indexers simply have nothing to read until a chain is created.
log "RayUp mode: skipping the Blockscout instance — no chain to index until you create one in the UI."

# Blockscout DB the worker indexer reads (override via remote.env if set).
#
# Deliberately left EMPTY in --rayup: ops-api gates its Blockscout indexers on this conn
# string, and when it IS set it verifies the chain id against BLOCKCHAIN_RPC_URL at boot
# and REFUSES TO START if that RPC is unreachable (di/container.go — the guard against
# ingesting another network's tokens). Chain-less there is no RPC to verify against, so
# setting it would kill ops-api on boot. Empty = indexers stay disabled until a chain
# exists, which is exactly what we want.

# ----------------------------------------------------------- bootstrap (bg job)
# Wait for ops-api to be healthy, create the admin via /admin/bootstrap, then grant
# the on-chain roles its wallet needs (login + token deploy). Mirrors the relayer's
# bootstrap_ops_api. Runs in the background so `compose up` can hold the foreground.
bootstrap_ops_api() {
    set +e
    if [ -z "$OPS_ADMIN_EMAIL" ]; then
        echo "[bootstrap] OPS_ADMIN_EMAIL not set — skipping admin bootstrap (set it in $REMOTE_ENV_FILE or pass --bootstrap-email)"
        return 0
    fi
    command -v jq >/dev/null 2>&1 || { echo "[bootstrap] jq missing — skipping"; return 0; }

    # Bootstrap targets the IDENTITY service: the users table lives there now, shared by
    # every chain, so there is exactly one admin account rather than one per chain.
    local url="http://localhost:8090" elapsed=0 code tmp addr role
    echo "[bootstrap] Waiting for identity /health…"
    while [ $elapsed -lt 300 ]; do
        curl -sf -o /dev/null --max-time 2 "$url/health" && break
        sleep 5; elapsed=$((elapsed + 5))
    done
    if [ $elapsed -ge 300 ]; then echo "[bootstrap] ! identity not healthy after 5 min — skipping"; return 0; fi

    echo "[bootstrap] POST $url/admin/bootstrap for $OPS_ADMIN_EMAIL"
    tmp="$(mktemp)"
    code="$(curl -sS -o "$tmp" -w '%{http_code}' -X POST "$url/admin/bootstrap" \
        -H 'Content-Type: application/json' -d "{\"email\":\"${OPS_ADMIN_EMAIL}\"}" 2>/dev/null)"
    [ -z "$code" ] && code="000"

    case "$code" in
        201)
            addr="$(jq -r '.address' "$tmp" 2>/dev/null)"
            echo "[bootstrap] admin created, wallet: $addr"
            ;;
        409)
            # Already bootstrapped — resolve the admin wallet from shared-db so we can
            # (idempotently) ensure its roles even when no fresh bootstrap happened.
            echo "[bootstrap] already bootstrapped — resolving admin wallet to ensure roles…"
            # ops_identity, not ops_api: users/user_wallets moved to the identity database.
            addr="$(docker exec shared-db psql -U blockscout -d ops_identity -At \
                -c "select w.rayls_address from users u join user_wallets w on w.user_id=u.id where u.email='${OPS_ADMIN_EMAIL}' limit 1;" 2>/dev/null | tr -d '[:space:]')"
            ;;
        *)
            echo "[bootstrap] ! /admin/bootstrap returned HTTP $code:"; cat "$tmp"; echo
            ;;
    esac
    rm -f "$tmp"

    # In this RayUp stack there is no chain (and no contracts) at boot, so on-chain role
    # grants happen per chain when RayUp provisions it. The admin user + custody wallet
    # created by /admin/bootstrap above is enough to authenticate.
    echo "[bootstrap] RayUp mode: admin created; on-chain role grants happen per chain."
}

# ----------------------------------------------------------------- up the stack
export PN_RPC_URL PN_RPC_URL_CONTAINER PN_CHAIN_ID DEPLOYMENT_PROXY_REGISTRY_ADDR="$REGISTRY_ADDR"
export PN_EXTRA_HOST WALLETCONNECT_PROJECT_ID BLOCKSCOUT_DB_CONN
export BLOCKCHAIN_STARTING_BLOCK="${STARTING_BLOCK:-0}"
# The ops database for the chain this run is bound to (the non-rayup modes). Empty = the
# compose default (ops_api). In rayup mode nothing sets it: each chain's own ops-api uses
# its own database (ops_api_<slug>), created and dropped with the chain by RayUp.
export OPS_DB_NAME="$(read_env OPS_DB_NAME "$REMOTE_ENV_FILE")"
export GOOGLE_CLIENT_ID="${GOOGLE_CLIENT_ID:-}" GOOGLE_CLIENT_SECRET="${GOOGLE_CLIENT_SECRET:-}"
export MICROSOFT_CLIENT_ID="${MICROSOFT_CLIENT_ID:-}" MICROSOFT_CLIENT_SECRET="${MICROSOFT_CLIENT_SECRET:-}"
# The explorer the BROWSER can actually open. In the bound-chain modes the chain is indexed
# by the local Blockscout instance resolved above ($BS_INSTANCE / $BS_HTTP_PORT), published
# on the host — so the reachable URL is always localhost:<that port>. Without this the
# playground has no real endpoint and falls back to the localStorage mock's slug template
# (https://<slug>.privatechain.rayls.com/explorer), a placeholder domain that does not
# resolve. Left empty in --rayup: there each chain gets its own explorer from the control
# plane, and the instance carries its own explorerUrl.
# Chain-less at boot: each RayUp chain carries its own explorer + RPC (the instance holds
# its explorerUrl), so these are empty here.
export NEXT_PUBLIC_EXPLORER_URL=""
export NEXT_PUBLIC_CHAIN_RPC_URL_PUBLIC=""

# Chain label shown in the playground (compose interpolates it).
#
# NOTE the playground's getDeploymentChain() classifies the instance by looking for
# "public" in this string — so a privacy chain's label must never contain that word,
# or the stablecoin wizard greys out the wrong deploy target.
# A RayUp axyl chain IS a privacy node (and is feeless, like the DEV PNs — no faucet).
export CHAIN_DISPLAY_NAME="Privacy Node (RayUp)"
export BLOCKCHAIN_FAUCET_PRIVATE_KEY=""
# Chain-less, permanently: the host stack is never bound to a chain. Every chain gets its
# OWN ops-api from RayUp, with that chain's config.
export CHAINLESS="true"
# Let the playground manage instances against the RayUp control plane (its /api/rayup/*
# proxy dereferences these SERVER-side; the browser only sees the same-origin proxy, never
# the token). The RayUp API runs in RayUp's OWN compose project, published on the host —
# reached from inside the playground container via host.docker.internal.
export RAYUP_API_URL_PLAYGROUND="http://host.docker.internal:4000"
export RAYUP_API_TOKEN
# Per-user chain scoping: the playground proxy forwards the caller's email with this token
# so each account sees only its own chains. Matches the control plane above.
export RAYUP_SERVICE_TOKEN

PROFILE_ARGS=()
[ "$WITH_PLAYGROUND" = true ] && PROFILE_ARGS+=(--profile playground)

bootstrap_ops_api &
BOOTSTRAP_PID=$!
BG_PIDS="$BOOTSTRAP_PID"

# No selection watcher in rayup mode any more: a chain created from the UI arrives with its
# OWN ops stack (RayUp provisions ops-<slug> alongside it), so there is nothing here to bind
# it to and nothing to watch for.

# shellcheck disable=SC2086  # word-splitting is intended: BG_PIDS is a pid list
trap 'kill $BG_PIDS 2>/dev/null || true' EXIT

# No ops-api/worker here: ops is per chain, provisioned with it by RayUp.
log "Starting the stack (identity, custody, postgres, nats$([ "$WITH_PLAYGROUND" = true ] && echo ', playground'))…"
log "  identity:   http://localhost:8090"
[ "$WITH_PLAYGROUND" = true ] && log "  playground: http://localhost:3000"
log "  chain:      none yet — create one in the playground (\"Create Sovereign Chain\")."
log "              RayUp deploys its contracts AND its own ops-api + explorer; the"
log "              playground routes to that chain's ops-api automatically."
log "  RayUp API:  $RAYUP_API_URL   (wipe + restart with: ./start_dev.sh --clean)"
# Hot-reload comes from air + the bind-mounted source (.:/app), so no compose --watch.
compose "${PROFILE_ARGS[@]}" up --build
