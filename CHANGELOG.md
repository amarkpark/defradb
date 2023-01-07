<a name="v0.4.0"></a>
## [v0.4.0](https://github.com/sourcenetwork/defradb/compare/v0.3.1...v0.4.0)

> 2023-12-23

DefraDB v0.4 is a major pre-production release. Until the stable version 1.0 is reached, the SemVer minor patch number will denote notable releases, which will give the project freedom to experiment and explore potentially breaking changes.

There are various new features in this release - some of which are breaking - and we invite you to review the official changelog below. Some highlights are persistence of replicators, DateTime scalars, TLS support, and GQL subscriptions.

This release does include a Breaking Change to existing v0.3.x databases. If you need help migrating an existing deployment, reach out at [hello@source.network](mailto:hello@source.network) or join our Discord at https://discord.source.network/.

### Features

* Add basic metric functionality ([#971](https://github.com/sourcenetwork/defradb/issues/971))
* Add thread safe transactional in-memory datastore ([#947](https://github.com/sourcenetwork/defradb/issues/947))
* Persist p2p replicators ([#960](https://github.com/sourcenetwork/defradb/issues/960))
* Add DateTime custom scalars ([#931](https://github.com/sourcenetwork/defradb/issues/931))
* Add GraphQL subscriptions ([#934](https://github.com/sourcenetwork/defradb/issues/934))
* Add support for tls ([#885](https://github.com/sourcenetwork/defradb/issues/885))
* Add group by support for commits ([#887](https://github.com/sourcenetwork/defradb/issues/887))
* Add depth support for commits ([#889](https://github.com/sourcenetwork/defradb/issues/889))
* Make dockey optional for allCommits queries ([#847](https://github.com/sourcenetwork/defradb/issues/847))
* Add WithStack to the errors package ([#870](https://github.com/sourcenetwork/defradb/issues/870))
* Add event system ([#834](https://github.com/sourcenetwork/defradb/issues/834))

### Fixes

* Correct errors.WithStack behaviour ([#984](https://github.com/sourcenetwork/defradb/issues/984))
* Correctly handle nested one to one joins ([#964](https://github.com/sourcenetwork/defradb/issues/964))
* Do not assume parent record exists when joining ([#963](https://github.com/sourcenetwork/defradb/issues/963))
* Change time format for HTTP API log ([#910](https://github.com/sourcenetwork/defradb/issues/910))
* Error if group select contains non-group-by fields ([#898](https://github.com/sourcenetwork/defradb/issues/898))
* Add inspection of values for ENV flags ([#900](https://github.com/sourcenetwork/defradb/issues/900))
* Remove panics from document ([#881](https://github.com/sourcenetwork/defradb/issues/881))
* Add __typename support ([#871](https://github.com/sourcenetwork/defradb/issues/871))
* Handle subscriber close ([#877](https://github.com/sourcenetwork/defradb/issues/877))
* Publish update events post commit ([#866](https://github.com/sourcenetwork/defradb/issues/866))

### Refactoring

* Make rootstore require Batching and TxnDatastore ([#940](https://github.com/sourcenetwork/defradb/issues/940))
* Conceptually clarify schema vs query-language ([#924](https://github.com/sourcenetwork/defradb/issues/924))
* Decouple db.db from gql ([#912](https://github.com/sourcenetwork/defradb/issues/912))
* Merkle clock heads cleanup ([#918](https://github.com/sourcenetwork/defradb/issues/918))
* Simplify dag fetcher ([#913](https://github.com/sourcenetwork/defradb/issues/913))
* Cleanup parsing logic ([#909](https://github.com/sourcenetwork/defradb/issues/909))
* Move planner outside the gql directory ([#907](https://github.com/sourcenetwork/defradb/issues/907))
* Refactor commit nodes ([#892](https://github.com/sourcenetwork/defradb/issues/892))
* Make latest commits syntax sugar ([#890](https://github.com/sourcenetwork/defradb/issues/890))
* Remove commit query ([#841](https://github.com/sourcenetwork/defradb/issues/841))

### Testing

* Add event tests ([#965](https://github.com/sourcenetwork/defradb/issues/965))
* Add new setup for testing explain functionality ([#949](https://github.com/sourcenetwork/defradb/issues/949))
* Add txn relation-type delete and create tests ([#875](https://github.com/sourcenetwork/defradb/issues/875))
* Skip change detection for tests that assert panic ([#883](https://github.com/sourcenetwork/defradb/issues/883))

### Continuous integration

* Bump all gh-action versions to support node16 ([#990](https://github.com/sourcenetwork/defradb/issues/990))
* Bump ssh-agent action to v0.7.0 ([#978](https://github.com/sourcenetwork/defradb/issues/978))
* Add error message format check ([#901](https://github.com/sourcenetwork/defradb/issues/901))

### Chore

* Extract (events, merkle) errors to errors.go ([#973](https://github.com/sourcenetwork/defradb/issues/973))
* Extract (datastore, db) errors to errors.go ([#969](https://github.com/sourcenetwork/defradb/issues/969))
* Extract (connor, crdt, core) errors to errors.go ([#968](https://github.com/sourcenetwork/defradb/issues/968))
* Extract inline (http and client) errors to errors.go ([#967](https://github.com/sourcenetwork/defradb/issues/967))
* Update badger version ([#966](https://github.com/sourcenetwork/defradb/issues/966))
* Move Option and Enumerable to immutables ([#939](https://github.com/sourcenetwork/defradb/issues/939))
* Add configuration of external loggers ([#942](https://github.com/sourcenetwork/defradb/issues/942))
* Strip DSKey prefixes and simplify NewDataStoreKey ([#944](https://github.com/sourcenetwork/defradb/issues/944))
* Include version metadata in cross-building ([#930](https://github.com/sourcenetwork/defradb/issues/930))
* Update to v0.23.2 the libP2P package ([#908](https://github.com/sourcenetwork/defradb/issues/908))
* Remove `ipfslite` dependency ([#739](https://github.com/sourcenetwork/defradb/issues/739))


<a name="v0.3.1"></a>
## [v0.3.1](https://github.com/sourcenetwork/defradb/compare/v0.3.0...v0.3.1)

> 2022-09-23

DefraDB v0.3.1 is a minor release, primarily focusing on additional/extended features and fixes of items added in the `v0.3.0` release.

### Features

* Add cid support for allCommits ([#857](https://github.com/sourcenetwork/defradb/issues/857))
* Add offset support to allCommits ([#859](https://github.com/sourcenetwork/defradb/issues/859))
* Add limit support to allCommits query ([#856](https://github.com/sourcenetwork/defradb/issues/856))
* Add order support to allCommits ([#845](https://github.com/sourcenetwork/defradb/issues/845))
* Display CLI usage on user error ([#819](https://github.com/sourcenetwork/defradb/issues/819))
* Add support for dockey filters in child joins ([#806](https://github.com/sourcenetwork/defradb/issues/806))
* Add sort support for numeric aggregates ([#786](https://github.com/sourcenetwork/defradb/issues/786))
* Allow filtering by nil ([#789](https://github.com/sourcenetwork/defradb/issues/789))
* Add aggregate offset support ([#778](https://github.com/sourcenetwork/defradb/issues/778))
* Remove filter depth limit ([#777](https://github.com/sourcenetwork/defradb/issues/777))
* Add support for and-or inline array aggregate filters ([#779](https://github.com/sourcenetwork/defradb/issues/779))
* Add limit support for aggregates ([#771](https://github.com/sourcenetwork/defradb/issues/771))
* Add support for inline arrays of nillable types ([#759](https://github.com/sourcenetwork/defradb/issues/759))
* Create errors package ([#548](https://github.com/sourcenetwork/defradb/issues/548))
* Add ability to display peer id ([#719](https://github.com/sourcenetwork/defradb/issues/719))
* Add a config option to set the vlog max file size ([#743](https://github.com/sourcenetwork/defradb/issues/743))
* Explain `topLevelNode` like a `MultiNode` plan ([#749](https://github.com/sourcenetwork/defradb/issues/749))
* Make `topLevelNode` explainable ([#737](https://github.com/sourcenetwork/defradb/issues/737))

### Fixes

* Order subtype without selecting the join child ([#810](https://github.com/sourcenetwork/defradb/issues/810))
* Correctly handles nil one-one joins ([#837](https://github.com/sourcenetwork/defradb/issues/837))
* Reset scan node for each join ([#828](https://github.com/sourcenetwork/defradb/issues/828))
* Handle filter input field argument being nil ([#787](https://github.com/sourcenetwork/defradb/issues/787))
* Ensure CLI outputs JSON to stdout when directed to pipe ([#804](https://github.com/sourcenetwork/defradb/issues/804))
* Error if given the wrong side of a one-one relationship ([#795](https://github.com/sourcenetwork/defradb/issues/795))
* Add object marker to enable return of empty docs ([#800](https://github.com/sourcenetwork/defradb/issues/800))
* Resolve the extra `typeIndexJoin`s for `_avg` aggregate ([#774](https://github.com/sourcenetwork/defradb/issues/774))
* Remove _like filter operator ([#797](https://github.com/sourcenetwork/defradb/issues/797))
* Remove having gql types ([#785](https://github.com/sourcenetwork/defradb/issues/785))
* Error if child _group selected without parent groupBy ([#781](https://github.com/sourcenetwork/defradb/issues/781))
* Error nicely on missing field specifier ([#782](https://github.com/sourcenetwork/defradb/issues/782))
* Handle order input field argument being nil ([#701](https://github.com/sourcenetwork/defradb/issues/701))
* Change output to outputpath in config file template for logger ([#716](https://github.com/sourcenetwork/defradb/issues/716))
* Delete mutations not correct persisting all keys ([#731](https://github.com/sourcenetwork/defradb/issues/731))

### Tooling

* Ban the usage of `ioutil` package ([#747](https://github.com/sourcenetwork/defradb/issues/747))
* Migrate from CircleCi to GitHub Actions ([#679](https://github.com/sourcenetwork/defradb/issues/679))

### Documentation

* Clarify meaning of url param, update in-repo CLI docs ([#814](https://github.com/sourcenetwork/defradb/issues/814))
* Disclaimer of exposed to network and not encrypted ([#793](https://github.com/sourcenetwork/defradb/issues/793))
* Update logo to respect theme ([#728](https://github.com/sourcenetwork/defradb/issues/728))

### Refactoring

* Replace all `interface{}` with `any` alias ([#805](https://github.com/sourcenetwork/defradb/issues/805))
* Use fastjson to parse mutation data string ([#772](https://github.com/sourcenetwork/defradb/issues/772))
* Rework limit node flow ([#767](https://github.com/sourcenetwork/defradb/issues/767))
* Make Option immutable ([#769](https://github.com/sourcenetwork/defradb/issues/769))
* Rework sum and count nodes to make use of generics ([#757](https://github.com/sourcenetwork/defradb/issues/757))
* Remove some possible panics from codebase ([#732](https://github.com/sourcenetwork/defradb/issues/732))
* Change logging calls to use feedback in CLI package ([#714](https://github.com/sourcenetwork/defradb/issues/714))

### Testing

* Add tests for aggs with nil filters ([#813](https://github.com/sourcenetwork/defradb/issues/813))
* Add not equals filter tests ([#798](https://github.com/sourcenetwork/defradb/issues/798))
* Fix `cli/peerid_test` to not clash addresses ([#766](https://github.com/sourcenetwork/defradb/issues/766))
* Add change detector summary to test readme ([#754](https://github.com/sourcenetwork/defradb/issues/754))
* Add tests for inline array grouping ([#752](https://github.com/sourcenetwork/defradb/issues/752))

### Continuous integration

* Reduce test resource usage and test with file db ([#791](https://github.com/sourcenetwork/defradb/issues/791))
* Add makefile target to verify the local module cache ([#775](https://github.com/sourcenetwork/defradb/issues/775))
* Allow PR titles to end with a number ([#745](https://github.com/sourcenetwork/defradb/issues/745))
* Add a workflow to validate pull request titles ([#734](https://github.com/sourcenetwork/defradb/issues/734))
* Fix the linter version to `v1.47` ([#726](https://github.com/sourcenetwork/defradb/issues/726))

### Chore

* Remove file system paths from resulting executable ([#831](https://github.com/sourcenetwork/defradb/issues/831))
* Add goimports linter for consistent imports ordering ([#816](https://github.com/sourcenetwork/defradb/issues/816))
* Improve UX by providing more information ([#802](https://github.com/sourcenetwork/defradb/issues/802))
* Change to defra errors and handle errors stacktrace ([#794](https://github.com/sourcenetwork/defradb/issues/794))
* Clean up `go.mod` with pruned module graphs ([#756](https://github.com/sourcenetwork/defradb/issues/756))
* Update to v0.20.3 of libp2p ([#740](https://github.com/sourcenetwork/defradb/issues/740))
* Bump to GoLang `v1.18` ([#721](https://github.com/sourcenetwork/defradb/issues/721))


<a name="v0.3.0"></a>
## [v0.3.0](https://github.com/sourcenetwork/defradb/compare/v0.2.1...v0.3.0)

> 2022-08-02

DefraDB v0.3 is a major pre-production release. Until the stable version 1.0 is reached, the SemVer minor patch number will denote notable releases, which will give the project freedom to experiment and explore potentially breaking changes.

There are *several* new features in this release, and we invite you to review the official changelog below. Some highlights are various new features for Grouping & Aggregation for the query system, like top-level aggregation and group filtering. Moreover, a brand new Query Explain system was added to introspect the execution plans created by DefraDB. Lastly we introduced a revamped CLI configuration system.

This release does include a Breaking Change to existing v0.2.x databases. If you need help migrating an existing deployment, reach out at [hello@source.network](mailto:hello@source.network) or join our Discord at https://discord.source.network/.

### Features

* Add named config overrides ([#659](https://github.com/sourcenetwork/defradb/issues/659))
* Expose color and caller log options, add validation ([#652](https://github.com/sourcenetwork/defradb/issues/652))
* Add ability to explain `groupNode` and it's attribute(s). ([#641](https://github.com/sourcenetwork/defradb/issues/641))
* Add primary directive for schema definitions ([@primary](https://github.com/primary)) ([#650](https://github.com/sourcenetwork/defradb/issues/650))
* Add support for aggregate filters on inline arrays ([#622](https://github.com/sourcenetwork/defradb/issues/622))
* Add explainable renderLimitNode & hardLimitNode attributes. ([#614](https://github.com/sourcenetwork/defradb/issues/614))
* Add support for top level aggregates ([#594](https://github.com/sourcenetwork/defradb/issues/594))
* Update `countNode` explanation to be consistent. ([#600](https://github.com/sourcenetwork/defradb/issues/600))
* Add support for stdin as input in CLI ([#608](https://github.com/sourcenetwork/defradb/issues/608))
* Explain `cid` & `field` attributes for `dagScanNode` ([#598](https://github.com/sourcenetwork/defradb/issues/598))
* Add ability to explain `dagScanNode` attribute(s). ([#560](https://github.com/sourcenetwork/defradb/issues/560))
* Add the ability to send user feedback to the console even when logging to file. ([#568](https://github.com/sourcenetwork/defradb/issues/568))
* Add ability to explain `sortNode` attribute(s). ([#558](https://github.com/sourcenetwork/defradb/issues/558))
* Add ability to explain `sumNode` attribute(s). ([#559](https://github.com/sourcenetwork/defradb/issues/559))
* Introduce top-level config package ([#389](https://github.com/sourcenetwork/defradb/issues/389))
* Add ability to explain `updateNode` attributes. ([#514](https://github.com/sourcenetwork/defradb/issues/514))
* Add `typeIndexJoin` explainable attributes. ([#499](https://github.com/sourcenetwork/defradb/issues/499))
* Add support to explain `countNode` attributes. ([#504](https://github.com/sourcenetwork/defradb/issues/504))
* Add CORS capability to HTTP API ([#467](https://github.com/sourcenetwork/defradb/issues/467))
* Add explaination of spans for `scanNode`. ([#492](https://github.com/sourcenetwork/defradb/issues/492))
* Add ability to Explain the response plan. ([#385](https://github.com/sourcenetwork/defradb/issues/385))
* Add aggregate filter support for groups only ([#426](https://github.com/sourcenetwork/defradb/issues/426))
* Configurable caller option in logger ([#416](https://github.com/sourcenetwork/defradb/issues/416))
* Add Average aggregate support ([#383](https://github.com/sourcenetwork/defradb/issues/383))
* Allow summation of aggregates ([#341](https://github.com/sourcenetwork/defradb/issues/341))
* Add ability to check DefraDB CLI version. ([#339](https://github.com/sourcenetwork/defradb/issues/339))

### Fixes

* Add a check to ensure limit is not 0 when evaluating query limit and offset ([#706](https://github.com/sourcenetwork/defradb/issues/706))
* Support multiple `--logger` flags ([#704](https://github.com/sourcenetwork/defradb/issues/704))
* Return without an error if relation is finalized ([#698](https://github.com/sourcenetwork/defradb/issues/698))
* Logger not correctly applying named config ([#696](https://github.com/sourcenetwork/defradb/issues/696))
* Add content-type media type parsing ([#678](https://github.com/sourcenetwork/defradb/issues/678))
* Remove portSyncLock deadlock condition ([#671](https://github.com/sourcenetwork/defradb/issues/671))
* Silence cobra default errors and usage printing ([#668](https://github.com/sourcenetwork/defradb/issues/668))
* Add stdout validation when setting logging output path ([#666](https://github.com/sourcenetwork/defradb/issues/666))
* Consider `--logoutput` CLI flag properly ([#645](https://github.com/sourcenetwork/defradb/issues/645))
* Handle errors and responses in CLI `client` commands ([#579](https://github.com/sourcenetwork/defradb/issues/579))
* Rename aggregate gql types ([#638](https://github.com/sourcenetwork/defradb/issues/638))
* Error when attempting to insert value into relationship field ([#632](https://github.com/sourcenetwork/defradb/issues/632))
* Allow adding of new schema to database ([#635](https://github.com/sourcenetwork/defradb/issues/635))
* Correctly parse dockey in broadcast log event. ([#631](https://github.com/sourcenetwork/defradb/issues/631))
* Increase system's open files limit in integration tests ([#627](https://github.com/sourcenetwork/defradb/issues/627))
* Avoid populating `order.ordering` with empties. ([#618](https://github.com/sourcenetwork/defradb/issues/618))
* Change to supporting of non-null inline arrays ([#609](https://github.com/sourcenetwork/defradb/issues/609))
* Assert fields exist in collection before saving to them ([#604](https://github.com/sourcenetwork/defradb/issues/604))
* CLI `init` command to reinitialize only config file ([#603](https://github.com/sourcenetwork/defradb/issues/603))
* Add config and registry clearing to TestLogWritesMessagesToFeedbackLog ([#596](https://github.com/sourcenetwork/defradb/issues/596))
* Change `$eq` to `_eq` in the failing test. ([#576](https://github.com/sourcenetwork/defradb/issues/576))
* Resolve failing HTTP API tests via cleanup ([#557](https://github.com/sourcenetwork/defradb/issues/557))
* Ensure Makefile compatibility with macOS ([#527](https://github.com/sourcenetwork/defradb/issues/527))
* Separate out iotas in their own blocks. ([#464](https://github.com/sourcenetwork/defradb/issues/464))
* Use x/cases for titling instead of strings to handle deprecation ([#457](https://github.com/sourcenetwork/defradb/issues/457))
* Handle limit and offset in sub groups ([#440](https://github.com/sourcenetwork/defradb/issues/440))
* Issue preventing DB from restarting with no records ([#437](https://github.com/sourcenetwork/defradb/issues/437))
* log serving HTTP API before goroutine blocks ([#358](https://github.com/sourcenetwork/defradb/issues/358))

### Testing

* Add integration testing for P2P. ([#655](https://github.com/sourcenetwork/defradb/issues/655))
* Fix formatting of tests with no extra brackets ([#643](https://github.com/sourcenetwork/defradb/issues/643))
* Add tests for `averageNode` explain. ([#639](https://github.com/sourcenetwork/defradb/issues/639))
* Add schema integration tests ([#628](https://github.com/sourcenetwork/defradb/issues/628))
* Add tests for default properties ([#611](https://github.com/sourcenetwork/defradb/issues/611))
* Specify which collection to update in test framework ([#601](https://github.com/sourcenetwork/defradb/issues/601))
* Add tests for grouping by undefined value ([#543](https://github.com/sourcenetwork/defradb/issues/543))
* Add test for querying undefined field ([#544](https://github.com/sourcenetwork/defradb/issues/544))
* Expand commit query tests ([#541](https://github.com/sourcenetwork/defradb/issues/541))
* Add cid (time-travel) query tests ([#539](https://github.com/sourcenetwork/defradb/issues/539))
* Restructure and expand filter tests ([#512](https://github.com/sourcenetwork/defradb/issues/512))
* Basic unit testing of `node` package ([#503](https://github.com/sourcenetwork/defradb/issues/503))
* Test filter in filter tests ([#473](https://github.com/sourcenetwork/defradb/issues/473))
* Add test for deletion of records in a relationship ([#329](https://github.com/sourcenetwork/defradb/issues/329))
* Benchmark transaction iteration ([#289](https://github.com/sourcenetwork/defradb/issues/289))

### Refactoring

* Improve CLI error handling and fix small issues ([#649](https://github.com/sourcenetwork/defradb/issues/649))
* Add top-level `version` package ([#583](https://github.com/sourcenetwork/defradb/issues/583))
* Remove extra log levels ([#634](https://github.com/sourcenetwork/defradb/issues/634))
* Change `sortNode` to `orderNode`. ([#591](https://github.com/sourcenetwork/defradb/issues/591))
* Rework update and delete node to remove secondary planner ([#571](https://github.com/sourcenetwork/defradb/issues/571))
* Trim imported connor package  ([#530](https://github.com/sourcenetwork/defradb/issues/530))
* Internal doc restructure ([#471](https://github.com/sourcenetwork/defradb/issues/471))
* Copy-paste connor fork into repo ([#567](https://github.com/sourcenetwork/defradb/issues/567))
* Add safety to the tests, add ability to catch stderr logs and add output path validation ([#552](https://github.com/sourcenetwork/defradb/issues/552))
* Change handler functions implementation and response formatting ([#498](https://github.com/sourcenetwork/defradb/issues/498))
* Improve the HTTP API implementation ([#382](https://github.com/sourcenetwork/defradb/issues/382))
* Use new logger in net/api ([#420](https://github.com/sourcenetwork/defradb/issues/420))
* Rename NewCidV1_SHA2_256 to mixedCaps ([#415](https://github.com/sourcenetwork/defradb/issues/415))
* Remove utils package ([#397](https://github.com/sourcenetwork/defradb/issues/397))
* Rework planNode Next and Value(s) function  ([#374](https://github.com/sourcenetwork/defradb/issues/374))
* Restructure aggregate query syntax ([#373](https://github.com/sourcenetwork/defradb/issues/373))
* Remove dead code from client package and document remaining ([#356](https://github.com/sourcenetwork/defradb/issues/356))
* Restructure datastore keys ([#316](https://github.com/sourcenetwork/defradb/issues/316))
* Add commits lost during github outage ([#303](https://github.com/sourcenetwork/defradb/issues/303))
* Move public members out of core and base packages ([#295](https://github.com/sourcenetwork/defradb/issues/295))
* Make db stuff internal/private ([#291](https://github.com/sourcenetwork/defradb/issues/291))
* Rework client.DB to ensure interface contains only public types ([#277](https://github.com/sourcenetwork/defradb/issues/277))
* Remove GetPrimaryIndexDocKey from collection interface ([#279](https://github.com/sourcenetwork/defradb/issues/279))
* Remove DataStoreKey from (public) dockey struct ([#278](https://github.com/sourcenetwork/defradb/issues/278))
* Renormalize to ensure consistent file line termination. ([#226](https://github.com/sourcenetwork/defradb/issues/226))
* Strongly typed key refactor ([#17](https://github.com/sourcenetwork/defradb/issues/17))

### Documentation

* Use permanent link to BSL license document ([#692](https://github.com/sourcenetwork/defradb/issues/692))
* README update v0.3.0 ([#646](https://github.com/sourcenetwork/defradb/issues/646))
* Improve code documentation ([#533](https://github.com/sourcenetwork/defradb/issues/533))
* Add CONTRIBUTING.md ([#531](https://github.com/sourcenetwork/defradb/issues/531))
* Add package level docs for logging lib ([#338](https://github.com/sourcenetwork/defradb/issues/338))

### Tooling

* Include all touched packages in code coverage ([#673](https://github.com/sourcenetwork/defradb/issues/673))
* Use `gotestsum` over `go test` ([#619](https://github.com/sourcenetwork/defradb/issues/619))
* Update Github pull request template ([#524](https://github.com/sourcenetwork/defradb/issues/524))
* Fix the cross-build script ([#460](https://github.com/sourcenetwork/defradb/issues/460))
* Add test coverage html output ([#466](https://github.com/sourcenetwork/defradb/issues/466))
* Add linter rule for `goconst`. ([#398](https://github.com/sourcenetwork/defradb/issues/398))
* Add github PR template. ([#394](https://github.com/sourcenetwork/defradb/issues/394))
* Disable auto-fixing linter issues by default ([#429](https://github.com/sourcenetwork/defradb/issues/429))
* Fix linting of empty `else` code blocks ([#402](https://github.com/sourcenetwork/defradb/issues/402))
* Add the `gofmt` linter rule. ([#405](https://github.com/sourcenetwork/defradb/issues/405))
* Cleanup linter config file ([#400](https://github.com/sourcenetwork/defradb/issues/400))
* Add linter rule for copyright headers ([#360](https://github.com/sourcenetwork/defradb/issues/360))
* Organize our config files and tooling. ([#336](https://github.com/sourcenetwork/defradb/issues/336))
* Limit line length to 100 characters (linter check) ([#224](https://github.com/sourcenetwork/defradb/issues/224))
* Ignore db/tests folder and the bench marks. ([#280](https://github.com/sourcenetwork/defradb/issues/280))

### Continuous Integration

* Fix circleci cache permission errors. ([#371](https://github.com/sourcenetwork/defradb/issues/371))
* Ban extra elses ([#366](https://github.com/sourcenetwork/defradb/issues/366))
* Fix change-detection to not fail when new tests are added. ([#333](https://github.com/sourcenetwork/defradb/issues/333))
* Update golang-ci linter and explicit go-setup to use v1.17 ([#331](https://github.com/sourcenetwork/defradb/issues/331))
* Comment the benchmarking result comparison to the PR ([#305](https://github.com/sourcenetwork/defradb/issues/305))
* Add benchmark performance comparisons ([#232](https://github.com/sourcenetwork/defradb/issues/232))
* Add caching / storing of bench report on default branch ([#290](https://github.com/sourcenetwork/defradb/issues/290))
* Ensure full-benchmarks are ran on a PR-merge. ([#282](https://github.com/sourcenetwork/defradb/issues/282))
* Add ability to control benchmarks by PR labels. ([#267](https://github.com/sourcenetwork/defradb/issues/267))

### Chore

* Update APL to refer to D2 Foundation ([#711](https://github.com/sourcenetwork/defradb/issues/711))
* Update gitignore to include `cmd` folders ([#617](https://github.com/sourcenetwork/defradb/issues/617))
* Enable random execution order of tests ([#554](https://github.com/sourcenetwork/defradb/issues/554))
* Enable linters exportloopref, nolintlint, whitespace ([#535](https://github.com/sourcenetwork/defradb/issues/535))
* Add utility for generation of man pages ([#493](https://github.com/sourcenetwork/defradb/issues/493))
* Add Dockerfile ([#517](https://github.com/sourcenetwork/defradb/issues/517))
* Enable errorlint linter ([#520](https://github.com/sourcenetwork/defradb/issues/520))
* Binaries in`cmd` folder, examples in `examples` folder ([#501](https://github.com/sourcenetwork/defradb/issues/501))
* Improve log outputs ([#506](https://github.com/sourcenetwork/defradb/issues/506))
* Move testing to top-level `tests` folder ([#446](https://github.com/sourcenetwork/defradb/issues/446))
* Update dependencies ([#450](https://github.com/sourcenetwork/defradb/issues/450))
* Update go-ipfs-blockstore and ipfs-lite ([#436](https://github.com/sourcenetwork/defradb/issues/436))
* Update libp2p dependency to v0.19 ([#424](https://github.com/sourcenetwork/defradb/issues/424))
* Update ioutil package to io / os packages. ([#376](https://github.com/sourcenetwork/defradb/issues/376))
* git ignore vscode ([#343](https://github.com/sourcenetwork/defradb/issues/343))
* Updated README.md contributors section ([#292](https://github.com/sourcenetwork/defradb/issues/292))
* Update changelog v0.2.1 ([#252](https://github.com/sourcenetwork/defradb/issues/252))


<a name="v0.2.1"></a>
## [v0.2.1](https://github.com/sourcenetwork/defradb/compare/v0.2.0...v0.2.1)

> 2022-03-04

### Features

* Add ability to delete multiple documents using filter ([#206](https://github.com/sourcenetwork/defradb/issues/206))
* Add ability to delete multiple documents, using multiple ids ([#196](https://github.com/sourcenetwork/defradb/issues/196))

### Fixes

* Concurrency control of Document using RWMutex ([#213](https://github.com/sourcenetwork/defradb/issues/213))
* Only log errors and above when benchmarking ([#261](https://github.com/sourcenetwork/defradb/issues/261))
* Handle proper type conversion on sort nodes ([#228](https://github.com/sourcenetwork/defradb/issues/228))
* Return empty array if no values found ([#223](https://github.com/sourcenetwork/defradb/issues/223))
* Close fetcher on error ([#210](https://github.com/sourcenetwork/defradb/issues/210))
* Installing binary using defradb name ([#190](https://github.com/sourcenetwork/defradb/issues/190))

### Tooling

* Add short benchmark runner option ([#263](https://github.com/sourcenetwork/defradb/issues/263))

### Documentation

* Add data format changes documentation folder ([#89](https://github.com/sourcenetwork/defradb/issues/89))
* Correcting typos ([#143](https://github.com/sourcenetwork/defradb/issues/143))
* Update generated CLI docs ([#208](https://github.com/sourcenetwork/defradb/issues/208))
* Updated readme with P2P section ([#220](https://github.com/sourcenetwork/defradb/issues/220))
* Update old or missing license headers ([#205](https://github.com/sourcenetwork/defradb/issues/205))
* Update git-chglog config and template ([#195](https://github.com/sourcenetwork/defradb/issues/195))

### Refactoring

* Introduction of logging system ([#67](https://github.com/sourcenetwork/defradb/issues/67))
* Restructure db/txn/multistore structures ([#199](https://github.com/sourcenetwork/defradb/issues/199))
* Initialize database in constructor ([#211](https://github.com/sourcenetwork/defradb/issues/211))
* Purge all println and ban it ([#253](https://github.com/sourcenetwork/defradb/issues/253))

### Testing

* Detect and force breaking filesystem changes to be documented ([#89](https://github.com/sourcenetwork/defradb/issues/89))
* Boost collection test coverage ([#183](https://github.com/sourcenetwork/defradb/issues/183))

### Continuous integration

* Combine the Lint and Benchmark workflows so that the benchmark job depends on the lint job in one workflow ([#209](https://github.com/sourcenetwork/defradb/issues/209))
* Add rule to only run benchmark if other check are successful ([#194](https://github.com/sourcenetwork/defradb/issues/194))
* Increase linter timeout ([#230](https://github.com/sourcenetwork/defradb/issues/230))

### Chore

* Remove commented out code ([#238](https://github.com/sourcenetwork/defradb/issues/238))
* Remove dead code from multi node ([#186](https://github.com/sourcenetwork/defradb/issues/186))


<a name="v0.2.0"></a>
## [v0.2.0](https://github.com/sourcenetwork/defradb/compare/v0.1.0...v0.2.0)

> 2022-02-07

DefraDB v0.2 is a major pre-production release. Until the stable version 1.0 is reached, the SemVer minor patch number will denote notable releases, which will give the project freedom to experiment and explore potentially breaking changes.

This release is jam-packed with new features and a small number of breaking changes. Read the full changelog for a detailed description. Most notable features include a new Peer-to-Peer (P2P) data synchronization system, an expanded query system to support GroupBy & Aggregate operations, and lastly TimeTraveling queries allowing to query previous states of a document.

Much more than just that has been added to ensure we're building reliable software expected of any database, such as expanded test & benchmark suites, automated bug detection, performance gains, and more.

This release does include a Breaking Change to existing v0.1 databases regarding the internal data model, which affects the "Content Identifiers" we use to generate DocKeys and VersionIDs. If you need help migrating an existing deployment, reach out at hello@source.network or join our Discord at https://discord.source.network.

### Features

* Added Peer-to-Peer networking data synchronization ([#177](https://github.com/sourcenetwork/defradb/issues/177))
* TimeTraveling (History Traversing) query engine and doc fetcher ([#59](https://github.com/sourcenetwork/defradb/issues/59))
* Add Document Deletion with a Key ([#150](https://github.com/sourcenetwork/defradb/issues/150))
* Add support for sum aggregate ([#121](https://github.com/sourcenetwork/defradb/issues/121))
* Add support for lwwr scalar arrays (full replace on update) ([#115](https://github.com/sourcenetwork/defradb/issues/115))
* Add count aggregate support ([#102](https://github.com/sourcenetwork/defradb/issues/102))
* Add support for named relationships ([#108](https://github.com/sourcenetwork/defradb/issues/108))
* Add multi doc key lookup support ([#76](https://github.com/sourcenetwork/defradb/issues/76))
* Add basic group by functionality ([#43](https://github.com/sourcenetwork/defradb/issues/43))
* Update datastore packages to allow use of context ([#48](https://github.com/sourcenetwork/defradb/issues/48))

### Bug fixes

* Only add join if aggregating child object collection ([#188](https://github.com/sourcenetwork/defradb/issues/188))
* Handle errors generated during input object thunks ([#123](https://github.com/sourcenetwork/defradb/issues/123))
* Remove new types from in-memory cache on generate error ([#122](https://github.com/sourcenetwork/defradb/issues/122))
* Support relationships where both fields have the same name ([#109](https://github.com/sourcenetwork/defradb/issues/109))
* Handle errors generated in fields thunk ([#66](https://github.com/sourcenetwork/defradb/issues/66))
* Ensure OperationDefinition case has at least one selection([#24](https://github.com/sourcenetwork/defradb/pull/24))
* Close datastore iterator on scan close ([#56](https://github.com/sourcenetwork/defradb/pull/56)) (resulted in a panic when using limit)
* Close superseded iterators before orphaning ([#56](https://github.com/sourcenetwork/defradb/pull/56)) (fixes a panic in the join code) 
* Move discard to after error check ([#88](https://github.com/sourcenetwork/defradb/pull/88)) (did result in panic if transaction creation fails)
* Check for nil iterator before closing document fetcher ([#108](https://github.com/sourcenetwork/defradb/pull/108))

### Tooling
* Added benchmark suite ([#160](https://github.com/sourcenetwork/defradb/issues/160))

### Documentation

* Correcting comment typos ([#142](https://github.com/sourcenetwork/defradb/issues/142))
* Correcting README typos ([#140](https://github.com/sourcenetwork/defradb/issues/140))

### Testing

* Add transaction integration tests ([#175](https://github.com/sourcenetwork/defradb/issues/175))
* Allow running of tests using badger-file as well as IM options ([#128](https://github.com/sourcenetwork/defradb/issues/128))
* Add test datastore selection support ([#88](https://github.com/sourcenetwork/defradb/issues/88))

### Refactoring

* Datatype modification protection ([#138](https://github.com/sourcenetwork/defradb/issues/138))
* Cleanup Linter Complaints and Setup Makefile ([#63](https://github.com/sourcenetwork/defradb/issues/63))
* Rework document rendering to avoid data duplication and mutation ([#68](https://github.com/sourcenetwork/defradb/issues/68))
* Remove dependency on concrete datastore implementations from db package ([#51](https://github.com/sourcenetwork/defradb/issues/51))
* Remove all `errors.Wrap` and update them with `fmt.Errorf`. ([#41](https://github.com/sourcenetwork/defradb/issues/41))
* Restructure integration tests to provide better visibility ([#15](https://github.com/sourcenetwork/defradb/pull/15))
* Remove schemaless code branches ([#23](https://github.com/sourcenetwork/defradb/pull/23))

### Performance
* Add badger multi scan support ([#85](https://github.com/sourcenetwork/defradb/pull/85))
* Add support for range spans ([#86](https://github.com/sourcenetwork/defradb/pull/86))

### Continous integration

* Use more accurate test coverage. ([#134](https://github.com/sourcenetwork/defradb/issues/134))
* Disable Codecov's Patch Check
* Make codcov less strict for now to unblock development ([#125](https://github.com/sourcenetwork/defradb/issues/125))
* Add codecov config file. ([#118](https://github.com/sourcenetwork/defradb/issues/118))
* Add workflow that runs a job on AWS EC2 instance. ([#110](https://github.com/sourcenetwork/defradb/issues/110))
* Add Code Test Coverage with CodeCov ([#116](https://github.com/sourcenetwork/defradb/issues/116))
* Integrate GitHub Action for golangci-lint Annotations ([#106](https://github.com/sourcenetwork/defradb/issues/106))
* Add Linter Check to CircleCi ([#92](https://github.com/sourcenetwork/defradb/issues/92))

### Chore

* Remove the S1038 rule of the gosimple linter. ([#129](https://github.com/sourcenetwork/defradb/issues/129))
* Update to badger v3, and use badger as default in memory store ([#56](https://github.com/sourcenetwork/defradb/issues/56))
* Make Cid versions consistent ([#57](https://github.com/sourcenetwork/defradb/issues/57))


<a name="v0.1.0"></a>
## v0.1.0

> 2021-03-15
