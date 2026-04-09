# SDK-Go P0 Fixes: omitempty + Error Fields

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Fix 53 NON_NULL fields silently dropped by `omitempty` and add error handling to 10 mutations that swallow failures.

**Architecture:** Change genqlient config from `optional: pointer` + `use_struct_references: true` to `optional: pointer_omitempty` + `use_struct_references: false`. This makes NON_NULL struct fields value types (always serialized) and nullable fields pointer+omitempty (omitted when nil). Then add `error { __typename }` to 10 mutation responses.

**Tech Stack:** Go 1.24, genqlient v0.8.1, Caido GraphQL API

**Working directory:** `/Users/mambrozkiewicz/Documents/sdk-go`

**Reference skill:** `/go` - use for Go patterns and conventions

---

### Task 1: Update genqlient.yaml

**Files:**
- Modify: `genqlient.yaml`

- [ ] **Step 1: Change optional and use_struct_references settings**

```yaml
schema: graphql/schema.graphql
operations:
  - graphql/operations/*.graphql
generated: graphql/generated.go
package: graphql
optional: pointer_omitempty
use_struct_references: false

bindings:
  DateTime:
    type: string
  Timestamp:
    type: int64
  Sensitive:
    type: string
  Token:
    type: string
  Upload:
    type: io.Reader
  JSON:
    type: encoding/json.RawMessage
  JsonObject:
    type: encoding/json.RawMessage
  JsonRaw:
    type: encoding/json.RawMessage
  Blob:
    type: string
  Duration:
    type: int64
  HTTPQL:
    type: string
  Image:
    type: string
  Alias:
    type: string
  Port:
    type: int
  Rank:
    type: string
  Snapshot:
    type: string
  Uri:
    type: string
  Url:
    type: string
  Version:
    type: string
  Binary:
    type: string
```

Key changes:
- `optional: pointer` -> `optional: pointer_omitempty` (nullable fields get pointer+omitempty, NON_NULL fields stay value types without omitempty)
- `use_struct_references: true` removed (NON_NULL struct fields become value types, not pointers)

- [ ] **Step 2: Commit**

```bash
git add genqlient.yaml
git commit -m "fix: change genqlient config to fix omitempty on NON_NULL fields"
```

---

### Task 2: Pull latest schema

**Files:**
- Modify: `graphql/schema.graphql`

- [ ] **Step 1: Pull schema from npm**

```bash
make schema
```

- [ ] **Step 2: Verify schema updated**

```bash
head -5 graphql/schema.graphql
```

- [ ] **Step 3: Commit**

```bash
git add graphql/schema.graphql
git commit -m "chore: update schema to latest from @caido/schema-proxy"
```

---

### Task 3: Add error fields to 10 mutations

**Files:**
- Modify: `graphql/operations/environment.graphql`
- Modify: `graphql/operations/project.graphql`
- Modify: `graphql/operations/workflow.graphql`
- Modify: `graphql/operations/plugin.graphql`
- Modify: `graphql/operations/task.graphql`

- [ ] **Step 1: Fix environment.graphql - SelectEnvironment and DeleteEnvironment**

Replace the two mutations at the bottom of the file:

```graphql
mutation SelectEnvironment($id: ID) {
  selectEnvironment(id: $id) {
    environment {
      id
      name
    }
    error {
      __typename
      ... on OtherUserError {
        code
      }
    }
  }
}

mutation DeleteEnvironment($id: ID!) {
  deleteEnvironment(id: $id) {
    deletedId
    error {
      __typename
      ... on OtherUserError {
        code
      }
    }
  }
}
```

- [ ] **Step 2: Fix project.graphql - RenameProject and DeleteProject**

Replace the two mutations at the bottom of the file:

```graphql
mutation RenameProject($id: ID!, $name: String!) {
  renameProject(id: $id, name: $name) {
    project {
      id
      name
    }
    error {
      __typename
      ... on OtherUserError {
        code
      }
    }
  }
}

mutation DeleteProject($id: ID!) {
  deleteProject(id: $id) {
    deletedId
    error {
      __typename
      ... on OtherUserError {
        code
      }
    }
  }
}
```

- [ ] **Step 3: Fix workflow.graphql - RenameWorkflow, DeleteWorkflow, GlobalizeWorkflow, LocalizeWorkflow**

Replace these four mutations (keep RunActiveWorkflow, RunConvertWorkflow, ToggleWorkflow unchanged - they already have error fields):

```graphql
mutation RenameWorkflow($id: ID!, $name: String!) {
  renameWorkflow(id: $id, name: $name) {
    workflow {
      id
      name
    }
    error {
      __typename
      ... on OtherUserError {
        code
      }
    }
  }
}

mutation DeleteWorkflow($id: ID!) {
  deleteWorkflow(id: $id) {
    deletedId
    error {
      __typename
      ... on OtherUserError {
        code
      }
    }
  }
}

mutation GlobalizeWorkflow($id: ID!) {
  globalizeWorkflow(id: $id) {
    workflow {
      id
      name
      global
    }
    error {
      __typename
      ... on OtherUserError {
        code
      }
    }
  }
}

mutation LocalizeWorkflow($id: ID!) {
  localizeWorkflow(id: $id) {
    workflow {
      id
      name
      global
    }
    error {
      __typename
      ... on OtherUserError {
        code
      }
    }
  }
}
```

- [ ] **Step 4: Fix plugin.graphql - TogglePlugin**

Replace the TogglePlugin mutation:

```graphql
mutation TogglePlugin($id: ID!, $enabled: Boolean!) {
  togglePlugin(id: $id, enabled: $enabled) {
    plugin {
      id
      enabled
    }
    error {
      __typename
      ... on OtherUserError {
        code
      }
    }
  }
}
```

- [ ] **Step 5: Fix task.graphql - CancelTask**

Replace the CancelTask mutation:

```graphql
mutation CancelTask($id: ID!) {
  cancelTask(id: $id) {
    cancelledId
    error {
      __typename
      ... on OtherUserError {
        code
      }
    }
  }
}
```

- [ ] **Step 6: Commit**

```bash
git add graphql/operations/*.graphql
git commit -m "fix: add error fields to 10 mutations that silently swallowed failures"
```

---

### Task 4: Regenerate and fix compilation

**Files:**
- Modify: `graphql/generated.go` (auto-generated)
- Modify: all SDK wrapper files (`*.go` in root) that reference changed pointer types

- [ ] **Step 1: Regenerate**

```bash
make generate
```

- [ ] **Step 2: Attempt build to find type errors**

```bash
go build ./... 2>&1
```

The config change from `use_struct_references: true` to `false` will change NON_NULL struct fields from `*Type` to `Type` (pointer to value). SDK wrapper files that dereference these fields or pass them as pointers will fail to compile.

- [ ] **Step 3: Fix each compilation error**

For each SDK wrapper file that fails, the fix pattern is:
- If the code was `input.SomeField = &gen.SomeType{...}`, change to `input.SomeField = gen.SomeType{...}` (drop the `&`)
- If the code was `if input.SomeField != nil`, change to check for zero value or remove the check
- If the code accesses `payload.Error`, add nil checks matching the existing pattern in files that already handle errors (e.g., `CreateTamperRule` in tamper.go)

Common files that will need fixes (check each):
- `replay.go` - `StartReplayTaskInput.Connection`, `StartReplayTaskInput.Settings` changed from `*Type` to `Type`
- `intercept.go` - `InterceptOptionsInput` fields
- `tamper.go` - tamper section/operation input types

For each file, read the compilation error, find the line, and change `*Type` to `Type` or `&Type{}` to `Type{}`.

- [ ] **Step 4: Verify clean build**

```bash
go build ./... && go vet ./...
```

- [ ] **Step 5: Verify omitempty fix**

Spot-check the generated code to confirm NON_NULL fields no longer have omitempty:

```bash
grep 'Placeholders.*omitempty' graphql/generated.go
grep 'connection.*omitempty' graphql/generated.go
```

Both should return zero results. If they still have omitempty, the genqlient config change didn't work and needs debugging.

- [ ] **Step 6: Verify error fields present**

```bash
grep -c 'SelectEnvironmentError\|DeleteEnvironmentError\|RenameProjectError\|DeleteProjectError' graphql/generated.go
```

Should return non-zero counts for the new error union types.

- [ ] **Step 7: Commit**

```bash
git add -A
git commit -m "feat: regenerate with fixed omitempty and error fields"
```

---

### Task 5: Tag and release v0.4.0

- [ ] **Step 1: Verify everything builds clean**

```bash
go build ./... && go vet ./... && go test -race ./...
```

- [ ] **Step 2: Tag the release**

Switch to the correct GitHub account first:

```bash
gh-context cotton
```

Then tag:

```bash
git tag v0.4.0
git push origin main --tags
```

---

### Task 6: Update caido-mcp-server to use sdk-go v0.4.0

**Working directory:** `/Users/mambrozkiewicz/Documents/Caido-Repo`

**Files:**
- Modify: `go.mod`
- Modify: `internal/tools/send_request.go` (remove workaround)
- Modify: `internal/tools/select_environment.go` (remove workaround)
- Modify: `internal/tools/delete_findings.go` (remove workaround)
- Modify: `internal/tools/export_findings.go` (remove workaround)

- [ ] **Step 1: Update dependency**

```bash
go get github.com/caido-community/sdk-go@v0.4.0
go mod tidy
```

- [ ] **Step 2: Fix compilation errors from type changes**

The sdk-go v0.4.0 changes NON_NULL struct fields from pointer to value types. Fix each:
- `send_request.go`: Can now revert to using `client.Replay.SendRequest()` directly since the SDK's `ReplayEntrySettingsInput.Placeholders` no longer has omitempty. Remove all the `replaySettingsNoOmit`, `replayTaskInputNoOmit`, `startReplayTaskVars`, `startReplayTaskPayload`, `startReplayTaskResp`, `startReplayTaskMutation`, and `startReplayTask()` function. Replace with the SDK call using value types (not pointers) for `ConnectionInfoInput` and `ReplayEntrySettingsInput`.
- `select_environment.go`: Can now revert to using `client.Environments.Select()` since the SDK mutation now queries the error field. Remove all the `selectEnvVars`, `selectEnvPayload`, `selectEnvResp`, `selectEnvironmentMutation`, and `selectEnvironmentRaw()` function. The SDK's response type now includes an Error field.
- `delete_findings.go`: The oneof issue was caused by `Ids []string` without omitempty. With `pointer_omitempty`, nullable fields get omitempty. Check if the generated `DeleteFindingsInput.Ids` field now has proper tags. If the oneof is fixed in generated code, remove the `deleteFindingsByIDsVars`, etc. workaround. If not (slices may not change), keep the workaround.
- `export_findings.go`: Same as delete_findings - check if oneof is fixed, remove workaround if so.

- [ ] **Step 3: Build and verify**

```bash
go build -o caido-mcp-server ./cmd/mcp/ && go vet ./...
```

- [ ] **Step 4: Deploy**

```bash
cp caido-mcp-server mcp
```

- [ ] **Step 5: Commit**

```bash
git add go.mod go.sum internal/tools/
git commit -m "feat: upgrade sdk-go to v0.4.0, remove GraphQL workarounds"
```

---

### Task 7: Integration test via Pentest2

- [ ] **Step 1: Send refresh to Pentest2 tmux pane**

```bash
tmux send-keys -t '0:Pentest2.0' 'sdk-go v0.4.0 deployed. run /mcp to refresh and retest all 34 tools.' Enter
```

- [ ] **Step 2: Monitor results**

Check Pentest2 pane for test results. All previously failing tools should now pass:
- `send_request` - PASS (omitempty fixed upstream)
- `create_tamper_rule` - PASS (section field fixed)
- `select_environment` - PASS (error field added upstream)
- `delete_findings` by reporter - PASS (oneof workaround or upstream fix)
- `export_findings` by reporter - PASS (oneof workaround or upstream fix)
