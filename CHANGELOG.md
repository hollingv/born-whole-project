# Changelog
All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

- - -
## v0.10.0 - 2026-09-09
#### Features
- (**seo**) add faqs for SEO optimization - (8e3c45d) - hollingv

- - -

## v0.9.0 - 2026-09-09
#### Features
- (**kb**) update intaction circumcision facts - (b6508d0) - hollingv

- - -

## v0.8.0 - 2026-09-09
#### Features
- (**ui**) add AI prompt to Resources page - (ccfe5ec) - hollingv

- - -

## v0.7.0 - 2026-09-09
#### Features
- (**Makefile**) use the SITE_NAME extracted from config.go - (cc12399) - hollingv
- (**build**) place all generated files into TARGET_DIR=dist - (8ddc5e1) - hollingv
- (**cli**) rename cli tool and all directories to 'bw' - (2872b42) - hollingv
- (**cli**) rename gacli to bictl - (52f90e0) - hollingv
- (**kb**) update the kb with 'bictl harvest' - (81acfe0) - hollingv
- (**kb**) update kb for intaction and nocirc - (31154b4) - hollingv
- (**kb**) harvest latest info from remote sites - (d19aad6) - hollingv
- (**organizations**) add thumbnail images to org page - (9bfbc33) - hollingv
- (**ui**) relocate version string to below brand name on header - (4f94166) - hollingv
- (**ui**) organization cards have buttons along bottom - (a5118f4) - hollingv
- (**ui**) Use clearer 'genital cutting' language - (941f598) - hollingv
- (**ui**) change verbiage around genital cutting - (9822bab) - hollingv
- (**ui**) add drop down menu for parents and intactivists - (ddab487) - hollingv
- (**ui**) collapse mission and about pages - (6af24d1) - hollingv
- (**ui**) add IntactGlobal thumbnail image - (4dfbc2b) - hollingv
- (**ui**) add the brand-logo to the nav bar - (78c8fec) - hollingv
- (**ui**) render the Sources one per line - (d53e63d) - hollingv
- (**ui**) place footer always at the bottom - (413281b) - hollingv
- (**ui**) increase font on the header links - (97b4dd5) - hollingv
#### Bug Fixes
- (**kbextract**) add tags for proper extraction of legal kb - (81987ad) - hollingv
- (**organization**) adjust Intact Global news to be /press - (3f96416) - hollingv
- (**ui**) remove example question about choice - (9fcaabc) - hollingv
#### Tests
- add unit tests to eliminate no-test-file warnings - (1063e3e) - hollingv
#### Refactoring
- change 'bodily integrity commons' to 'born whole' - (5c6bcad) - hollingv

- - -

## v0.6.0 - 2026-08-11
#### Features
- (**Makefile**) add 'whitepaper' target - (412d48b) - hollingv
- (**bictl**) add --source switch to the 'harvest' command - (351e085) - hollingv
- (**ui**) add hyperlinks to the sources returned by ai query - (a69acce) - hollingv
- (**ui**) remove version from hamburger drop-down - (cf5cd82) - hollingv
- (**ui**) create better footer for all pages with siteName - (f4de045) - hollingv
- (**ui**) site background is pale purple - (950fd83) - hollingv
- add utility script to create pdf from md - (5bb3d43) - hollingv
- change the kb-build command to harvest - (77e8d27) - hollingv
#### Documentation
- improved whitepaper for submission - (225cbc8) - hollingv
- add README-whitepaper.md - (84903c8) - hollingv
- add a project overview page - (c03526d) - hollingv
#### Refactoring
- (**kb.go**) rename to harvest.go and associated renames - (e5cc5ec) - hollingv
- (**packages**) top level dir is commands and helpers - (8e3ded8) - hollingv
- (**styles.css**) add --brand-size to :root in styles.css - (0b86dca) - hollingv
- organize and document the styles.css - (1dde4a0) - hollingv

- - -

## v0.5.0 - 2026-08-09
#### Features
- (**Resources**) add sections, toc and 'top' button - (d940b8c) - hollingv
- (**kb**) add concurrency to the kb-build command - (2a2ffe6) - hollingv
- (**ui**) add toc and 'top' button on organizations.html.tmpl - (823c9ca) - hollingv
- (**ui**) consistent white color on all pages - (b046eb9) - hollingv
- add 'status' command to show state of all needed env vars - (a8178ca) - hollingv
- add youtube shorts to resources page - (fe3608f) - hollingv
- introduce 'how to submit' your organization - (229327a) - hollingv
#### Bug Fixes
- (**ci.yaml**) replace problematic projectPrefix - (ca213e1) - hollingv
- (**ci.yaml**) ensure 3 env vars are avail to all steps - (bc46cd1) - hollingv
#### Documentation
- (**index**) fix wording about consent - (391aa12) - hollingv
#### Tests
- add check of resources page - (f3e1a04) - hollingv
#### Refactoring
- (**ask.js**) extract functions to simplify - (e293f73) - hollingv
- (**ci.yaml**) simplify the deployment steps - (617d194) - hollingv
- (**kb.go**) create new package for helper functions - (0f9efa4) - hollingv
- extract server.go into new files in the main package - (6b54a72) - hollingv
- remove dead code and add comments and refactorings - (a015f26) - hollingv
- introduce projectPrefix as constants and rename env vars - (712b9c3) - hollingv

- - -

## v0.4.0 - 2026-08-07
#### Features
- (**UI**) make hero image smaller by 20% - (1e013b1) - hollingv
- (**ui**) add organization placeholder image to org card - (e48c8dd) - hollingv
- improve kb extraction quality by targeting content elements only - (34a8b2e) - hollingv
#### Bug Fixes
- adjust the test to correct ai mistake - (6c09db1) - hollingv
#### Documentation
- begin use of Bodily Integrity Commons naming - (2175bc9) - hollingv
#### Continuous Integration
- run 'make test' in pre-push hook to block broken pushes - (52aa0e8) - hollingv

- - -

## v0.3.1 - 2026-08-06
#### Bug Fixes
- (**ci.yaml**) ensure tags are fetched even if cached - (b08e90a) - hollingv
- cog toml should have tag_prefix = 'v' - (ee4c7ae) - hollingv
- ignore merge commits running cog - (62df321) - hollingv
- adjust hooks and cog check in Makefile - (4bfaa81) - hollingv

- - -

## [v0.3.0](https://github.com/IntactGlobal/ig_cli/compare/24762f81af8a84092746e094f2f697651b959a45..v0.3.0) - 2026-08-05
#### Features
- (**ui**) add revolving sample questions to the AI Ask - ([c0338b1](https://github.com/IntactGlobal/ig_cli/commit/c0338b1ef0cdefedf72cbdaf528d637774b079ef)) - hollingv
- improve kb search quality with paragraph-level chunking - ([043b6f2](https://github.com/IntactGlobal/ig_cli/commit/043b6f283567e91d9706f038be20ab5f65297d4e)) - hollingv
- home page Ask prompt is always on the right - ([f1a5e62](https://github.com/IntactGlobal/ig_cli/commit/f1a5e6235a498b6cd4c45bb541d3b054687842ef)) - hollingv
- add hyperlinked SiteName on nav bar - ([986de6c](https://github.com/IntactGlobal/ig_cli/commit/986de6c6aa0eb248234af9bec90b6f798234090d)) - hollingv
#### Documentation
- add readme - ([db83f5f](https://github.com/IntactGlobal/ig_cli/commit/db83f5fe27d6c374949dbc11a7e1d72723fcd651)) - hollingv
- ask.js and server.go should mirror each other - ([3006ab1](https://github.com/IntactGlobal/ig_cli/commit/3006ab1047e579a3a976c3b56da5f787c213cf17)) - hollingv
#### Refactoring
- remove duplicated server*.go files - ([24762f8](https://github.com/IntactGlobal/ig_cli/commit/24762f81af8a84092746e094f2f697651b959a45)) - hollingv

- - -

## [v0.2.0](https://github.com/IntactGlobal/ig_cli/compare/8f3fcad815023f15c56c2a77ad3fca02bf072379..v0.2.0) - 2026-08-04
#### Features
- (**ask.js**) catch Ask/submit error messages from Cloudflare - ([bfbe8de](https://github.com/IntactGlobal/ig_cli/commit/bfbe8de0a8fe29947720532fba8fcae943b429c6)) - hollingv
- increase font on Ask/submit response - ([49eb987](https://github.com/IntactGlobal/ig_cli/commit/49eb987c0f85cdc93f386ada547a221c7a0c6ed4)) - hollingv
- add stronger prompt wording and add disclaimer - ([9068804](https://github.com/IntactGlobal/ig_cli/commit/9068804d8d320237f31ba65e5906736d2422fa8b)) - hollingv
- update AI model to @cf/meta/llama-3.1-8b-instruct-fast - ([5d35ce3](https://github.com/IntactGlobal/ig_cli/commit/5d35ce39c495462be1ba29c950717aa2d769fac5)) - hollingv
- add keyword extraction and scoring to Ask search - ([a45e80d](https://github.com/IntactGlobal/ig_cli/commit/a45e80db85d671033e7b27e1022443056b2a863b)) - hollingv
- enable raw kb responses via cloudflare api calls - ([ac2a870](https://github.com/IntactGlobal/ig_cli/commit/ac2a87037293d754aaee4fa8d046ce0efc892d14)) - hollingv
- change "Visit Site" to be blue like other buttons - ([3fe3fdf](https://github.com/IntactGlobal/ig_cli/commit/3fe3fdfadc3040c4956726c3db84a34ad955b051)) - hollingv
- add new KBURLs based on suggestions from Claude Sonnet - ([2e359db](https://github.com/IntactGlobal/ig_cli/commit/2e359db161cf4d7a358865d33e3939113529f219)) - hollingv
- add faq page with 3 questions - ([685d002](https://github.com/IntactGlobal/ig_cli/commit/685d00201af845a986ccb8d9aa63ed799a22f3f3)) - hollingv
- add html table to Ask response with source - ([2fdf898](https://github.com/IntactGlobal/ig_cli/commit/2fdf8984b1387f877a06b1c61a77199ab2b04d0e)) - hollingv
- add list of KBURLs for each organization - ([2727c8a](https://github.com/IntactGlobal/ig_cli/commit/2727c8a3d74199e84c79b0ad8e05b2e47675f415)) - hollingv
- new knowledge build command returning site/kb for commit - ([f683f75](https://github.com/IntactGlobal/ig_cli/commit/f683f75aa8b97b3566be37581df900b1c10f0b94)) - hollingv
- create knowledge base with site url on 'gacli kb-build' - ([4b698a7](https://github.com/IntactGlobal/ig_cli/commit/4b698a7bbc87d68720f14709363162d0bf6bac64)) - hollingv
- add kb-build command as stubbed command - ([8f3fcad](https://github.com/IntactGlobal/ig_cli/commit/8f3fcad815023f15c56c2a77ad3fca02bf072379)) - hollingv
#### Bug Fixes
- use kb manifest to enable Ask on Cloudflare Pages - ([3607048](https://github.com/IntactGlobal/ig_cli/commit/36070485a80de82c9607046bf90c708e5a5f24d0)) - hollingv
- kill the server process after test target ends - ([6c1d57b](https://github.com/IntactGlobal/ig_cli/commit/6c1d57b72c2a6cdb0a09f26df482254d2d114c96)) - hollingv
#### Refactoring
- write site/kb/manifest.json in human-readable form - ([a616054](https://github.com/IntactGlobal/ig_cli/commit/a616054239e0b2f7d971a1d687a4d6ccf105d6e8)) - hollingv

- - -

## [v0.1.2](https://github.com/IntactGlobal/ig_cli/compare/e29c9eafaf99cd008b734878f7e9cf664824389e..v0.1.2) - 2026-08-03
#### Bug Fixes
- (**ci.yaml**) fetch full history including tags - ([e29c9ea](https://github.com/IntactGlobal/ig_cli/commit/e29c9eafaf99cd008b734878f7e9cf664824389e)) - hollingv

- - -

## [v0.1.1](https://github.com/IntactGlobal/ig_cli/compare/f49271ed40ebe1882498df660413648957c5077b..v0.1.1) - 2026-08-03
#### Bug Fixes
- (**Makefile**) add debug of APP_TAG - ([f49271e](https://github.com/IntactGlobal/ig_cli/commit/f49271ed40ebe1882498df660413648957c5077b)) - hollingv

- - -

## [v0.1.0](https://github.com/IntactGlobal/ig_cli/compare/aa545df65efcc6bf41a514634ae895c85112fa6d..v0.1.0) - 2026-08-03
#### Features
- each nav link is its own page - ([493ff17](https://github.com/IntactGlobal/ig_cli/commit/493ff174e611c6d01751009b92987a71bb7f2141)) - hollingv
- add Home nav bar link with tigher spacing - ([a22bd48](https://github.com/IntactGlobal/ig_cli/commit/a22bd48d36b9c5064e3bb0cd4ed4e82ce95fa503)) - hollingv
- nav bar links stay visible at half-size - ([4f1cb0c](https://github.com/IntactGlobal/ig_cli/commit/4f1cb0caf0397ed7f02dc6ae484e894fbb59b4d3)) - hollingv
- add a mission page and include on main navigation - ([78f8cad](https://github.com/IntactGlobal/ig_cli/commit/78f8cadaf6dbcf23d798fddf840ae891c57f85a7)) - hollingv
- add version to the hamburger list of options - ([4d3d52b](https://github.com/IntactGlobal/ig_cli/commit/4d3d52bf685ea04d8412739f585e0c58dc8065ba)) - hollingv
- introduce server side code and devel deployment - ([7ea5898](https://github.com/IntactGlobal/ig_cli/commit/7ea58985574f73541bcf60d8108bff03455977a4)) - hollingv
- reorder and correct the organization list - ([85ee681](https://github.com/IntactGlobal/ig_cli/commit/85ee68134f0a67ce9a7ace9744d1d9739f5737fa)) - hollingv
- add better card imagery and AI prompt form - ([04fa3be](https://github.com/IntactGlobal/ig_cli/commit/04fa3be81b44900892d9fc27214c068a7c5d6855)) - hollingv
- add blue-pink buttons for organizations - ([c6e892d](https://github.com/IntactGlobal/ig_cli/commit/c6e892dacef8eaef7347db97a2c974ffcc3c7ddc)) - hollingv
- imroved button for each organization - ([1ca17b1](https://github.com/IntactGlobal/ig_cli/commit/1ca17b1703daf0b45990598865f9535018cfaf45)) - hollingv
- improved UI with hamburger, const site and more orgs - ([21b41cb](https://github.com/IntactGlobal/ig_cli/commit/21b41cb2984b4961c3e9ef04aba3099628a217a9)) - hollingv
- html templating and css style sheet - ([4a4fe4d](https://github.com/IntactGlobal/ig_cli/commit/4a4fe4d8bdeec35ae2e84bcd2bf78a1157b886f7)) - hollingv
- introduce gacli and organization.yaml for future - ([b8b96e1](https://github.com/IntactGlobal/ig_cli/commit/b8b96e12ac04da6b3a30cf278f2cab7e0259cc3f)) - hollingv
- improve visual layout for large monitors - ([8946868](https://github.com/IntactGlobal/ig_cli/commit/894686800cd2e140ae3da8e59a6e4bdb0bb932f6)) - hollingv
- add much better hero image - ([58c9e68](https://github.com/IntactGlobal/ig_cli/commit/58c9e6824811430fdf2daa9760dc7b2c22fde98d)) - hollingv
- make site more modern looking on the index.html - ([a8492dd](https://github.com/IntactGlobal/ig_cli/commit/a8492dd103ff7cb7489899567564ce2bb26e483a)) - hollingv
#### Bug Fixes
- (**Makefile**) remove reference to publication - ([18d05cc](https://github.com/IntactGlobal/ig_cli/commit/18d05cc6eccdadba37553753e38bd44429bad505)) - hollingv
- (**integ_test.go**) look for orgs on organizations.html - ([8d2d9e7](https://github.com/IntactGlobal/ig_cli/commit/8d2d9e75cab357769613f8dec64e19a0b85ed3b0)) - hollingv
- remove duplicated NOCIRC org - ([2f8ce90](https://github.com/IntactGlobal/ig_cli/commit/2f8ce901353681f9f11e0fdb0c4ce024cace60dd)) - hollingv
- integ tests wait for site up - ([5b9209b](https://github.com/IntactGlobal/ig_cli/commit/5b9209ba9ee433cc38316715e611faf2da995b00)) - hollingv
- fixed links and css naming in addition to cards layout - ([5f2b043](https://github.com/IntactGlobal/ig_cli/commit/5f2b04374c402cbfbdfa22a44d105b3ce8281866)) - hollingv
- include the cmd/gacli dir - ([222cc55](https://github.com/IntactGlobal/ig_cli/commit/222cc5568f55596bdab8d1b23fa4af4790abc79f)) - hollingv
- reduce size of baby.jpg - ([aa545df](https://github.com/IntactGlobal/ig_cli/commit/aa545df65efcc6bf41a514634ae895c85112fa6d)) - hollingv
#### Documentation
- better wording on intact global - ([68053f9](https://github.com/IntactGlobal/ig_cli/commit/68053f9bd84fb18c093a766114274ab4eec92bdd)) - hollingv
- better wording on home page - ([d301499](https://github.com/IntactGlobal/ig_cli/commit/d3014990abca4d15b38647ba064ef4c63e9e7a0c)) - hollingv
- change title page - ([d869cf8](https://github.com/IntactGlobal/ig_cli/commit/d869cf8d529ec82ab875f15e3b4f7ac8198905d0)) - hollingv
#### Tests
- add unit and integration test infrastructure - ([cc6db33](https://github.com/IntactGlobal/ig_cli/commit/cc6db339c0b1309e589974b1e058e0009424baad)) - hollingv
#### Continuous Integration
- add branch name to the html footer - ([3fee3ea](https://github.com/IntactGlobal/ig_cli/commit/3fee3eadfbab92d368aadc164e6dbc39041c0c35)) - hollingv
- combine the build and deploy into 1 step - ([273cd2f](https://github.com/IntactGlobal/ig_cli/commit/273cd2f574006bbb9f37ffa6a59b4aa116161d41)) - hollingv
- inject the Makefile's APP_TAG into index.html - ([999308c](https://github.com/IntactGlobal/ig_cli/commit/999308c554ec8bacd9235058eb94cbbec20f65ca)) - hollingv
#### Refactoring
- query templates instead of hardcoding - ([2ab2a36](https://github.com/IntactGlobal/ig_cli/commit/2ab2a361e5c03d01e735940866004c8faa4221f3)) - hollingv
- rename the gacli 'html' to 'site' - ([673a84f](https://github.com/IntactGlobal/ig_cli/commit/673a84f49fd47d3f3f37970bc819307851a3b62e)) - hollingv
- remove dead ./content dir - ([1d1378b](https://github.com/IntactGlobal/ig_cli/commit/1d1378b41313f10603c8b1dd632498bab328d32b)) - hollingv
- use src directory for source code - ([fa8a0ff](https://github.com/IntactGlobal/ig_cli/commit/fa8a0ff9f6eeebbd837290eb36394374d9bac2ab)) - hollingv

- - -

Changelog generated by [cocogitto](https://github.com/cocogitto/cocogitto).