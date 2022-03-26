### Scope of the project (what specific, preferably one problem it solves)

Short-term (1-5 days) log aggregation from fluentd. Meant to be an deployed in a container alogside single project it serves. Thus write performance and memory footprint are critical.

### API schema design.

...

### Logging plan

...

### Monitoring plan.

...

### CI/CD:

- Go specific: golangci-lint
- env files linter
- precommit with bunch of meta rules (no whitespace, end of file, etc)
- your own scripts (check file line length limits)
- spell check

#### Test plan.

  - TDD. Unit test -> implementation -> refactoring. Make them fast. Mock external dependancies (database) if tests are slow.
  - Make time to set up integration/acceptance tests.
  - After project is operational, set up end-to-end testing for critical user journeys.
  - Performance testing (latency), load testing. (Don't do that just yet, or at all. That is probably premature optimization.)
  - If some parts are absolutely critical - set up mutation testing, fuzz testing (use something like schemathesis for APIs) and fault tolerance testing (if some dependancy is not available - for example, if database is down).

#### Backup plan (if-then)

Out of scope

