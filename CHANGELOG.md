#####

# Changelog

All notable changes to "gomisc" will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Go to [legend](#legend---types-of-changes) for further information about the types of changes.

## [Unreleased]

## [0.10.0] - 2023-09-04

### Added

- Package 'clipboard' and a usage example. [91ae789](https://github.com/Sven-Seyfert/gomisc/commit/91ae7899baed079e94b8aee21bc1f7369204299d)
- Tests for package 'clipboard'. [d466782](https://github.com/Sven-Seyfert/gomisc/commit/d466782fca04919877d1a341dcdfe7ae11eca52b)

### Changed

- Extend .gitignore file. [952b536](https://github.com/Sven-Seyfert/gomisc/commit/952b5366ac7de5951ff4def5372cebc41e7ffac7)
- Rename 'clipboard_window.go' to 'clipboard_windows.go'. [1efab72](https://github.com/Sven-Seyfert/gomisc/commit/1efab726294532432194760fc3257750946b05b6)
- Update coverage.html file. [425e0db](https://github.com/Sven-Seyfert/gomisc/commit/425e0db94cc7f412d0667c84bdd2f3e2a13f1e90)

### Documented

- Change badges in README.md file. [f24d29a](https://github.com/Sven-Seyfert/gomisc/commit/f24d29acac6216680f4f85c843bd947d8b20493b)
- Fix incorrect display of 'go coverage' badge. [775e639](https://github.com/Sven-Seyfert/gomisc/commit/775e639ff4c8ae88e0dc908df3543ad8aa9f5b74)
- Update 'go coverage' badge url. [df1967b](https://github.com/Sven-Seyfert/gomisc/commit/df1967be57c540190bc95982ff2eee12ea05dea3)
- Update CHANGELOG.md and README.md file. [47a8112](https://github.com/Sven-Seyfert/gomisc/commit/47a8112a24928f3731a520a4b11b1c9c79e1198a)

### Refactored

- Rename file cover.html to coverage.html. [7b48299](https://github.com/Sven-Seyfert/gomisc/commit/7b482992982199ed748b343ea7c96a2e35ecccb2)

## [0.9.0] - 2023-08-30

### Added

- Test coverage report file. [dd857a7](https://github.com/Sven-Seyfert/gomisc/commit/dd857a78f8293b178a00b0e461c4a38e528aadff)
- Tests for package 'crypt'. [85f40b7](https://github.com/Sven-Seyfert/gomisc/commit/85f40b7825e96072a198eb811d2290a561637f1f)

### Changed

- Error handling for package 'crypt'. [1d2a4cb](https://github.com/Sven-Seyfert/gomisc/commit/1d2a4cbd3c9d976558a8a5159e54683b3d9714bb)
- Extend tests for package 'singleinstance'. [36b2737](https://github.com/Sven-Seyfert/gomisc/commit/36b27377a7a9c9782c6cd78c4a850f9f1a0cd0d0)
- Update test coverage file. [9add680](https://github.com/Sven-Seyfert/gomisc/commit/9add6807985d1c2d4fac9fcb3a6671642f05aabf)

### Documented

- Update CHANGELOG.md and README.md file. [7b786ea](https://github.com/Sven-Seyfert/gomisc/commit/7b786ea8bf5cc33e5e1bb88880763dd64c89813b)

### Refactored

- Syntax adjustment of error message. [71dd730](https://github.com/Sven-Seyfert/gomisc/commit/71dd730681f91d7a569c801acf917c5ed154e30d)
- Use raw string (backtick) to avoid character escaping. [e556b6a](https://github.com/Sven-Seyfert/gomisc/commit/e556b6a499a7a7c0bed0f11dd7f0c78a93eb26e1)

## [0.8.0] - 2023-08-29

### Added

- Tests for package 'singleinstance'. [f558d76](https://github.com/Sven-Seyfert/gomisc/commit/f558d768a9612b3e81e21239beb16bc9a547ea84)

### Changed

- Error handling for package 'singleinstance'. [41712cd](https://github.com/Sven-Seyfert/gomisc/commit/41712cd8212b2dbdc963c997b788239fdcae48da)

### Documented

- Update CHANGELOG.md and README.md file. [31ec7f2](https://github.com/Sven-Seyfert/gomisc/commit/31ec7f2dfbc2fb4488ee9fefcd9f77906a68b087)

## [0.7.0] - 2023-08-25

### Changed

- Add .golangci.yml to .gitignore (linter settings). [0cb0150](https://github.com/Sven-Seyfert/gomisc/commit/0cb015072a1468524f002b1e3bcdc7c13c80c701)
- Apply linter suggestions. [b2d2550](https://github.com/Sven-Seyfert/gomisc/commit/b2d2550ddb585eae6a3ef325574ba44d09adf0c9)

### Documented

- Add package description for Godocs (pkg.go.dev). [3820f01](https://github.com/Sven-Seyfert/gomisc/commit/3820f01e0454ecc5aefc88c5d32d0d84b810b68d)
- Update CHANGELOG.md and README.md file. [b5c1943](https://github.com/Sven-Seyfert/gomisc/commit/b5c194357e7ca3302db5b1300c27878aae80bd62)

## [0.6.0] - 2023-08-23

### Documented

- Add missing description for package 'crypt'. [1853785](https://github.com/Sven-Seyfert/gomisc/commit/1853785a1a6cfc26cf47e40ae014bceced931f99)

## [0.5.0] - 2023-08-23

### Added

- Package 'crypt' and a usage example. [be48ee5](https://github.com/Sven-Seyfert/gomisc/commit/be48ee5fbeee954ccd089632ea6a24be3d03170d)

### Changed

- Extend example for 'singleinstance' package. [26316c2](https://github.com/Sven-Seyfert/gomisc/commit/26316c23883817dea31931e6375588def20f6cef)
- File structure of single_instance.go and documents functions by comments. [0f132eb](https://github.com/Sven-Seyfert/gomisc/commit/0f132ebb20909d302f20a2c589c1d9b21e271a95)

### Documented

- Update CHANGELOG.md and README.md file. [46a5bb0](https://github.com/Sven-Seyfert/gomisc/commit/46a5bb01b2b8d1a0836328fb60ce37850a664610)
- Update README.md file for introduced 'crypt' package. [ce31a23](https://github.com/Sven-Seyfert/gomisc/commit/ce31a2309c53ab387320a49e7283a2df583191b8)

## [0.4.0] - 2023-08-21

### Changed

- Extract error handling to separate function. [c4546bd](https://github.com/Sven-Seyfert/gomisc/commit/c4546bd19e7e3d377c53bf60c72be41a16bd0e83)

### Documented

- Update README.md by hint about purpose project. [906eaa4](https://github.com/Sven-Seyfert/gomisc/commit/906eaa4e135160f2fea3a449b4811215ae268611)

## [0.3.0] - 2023-08-20

### Changed

- My repository name from capital case to lower case. [59b4885](https://github.com/Sven-Seyfert/gomisc/commit/59b4885a3ea4a103ada07b0148cc3fbfb00b6f82)

### Documented

- Bring the original LICENSE file back. [b7eacff](https://github.com/Sven-Seyfert/gomisc/commit/b7eacffa56fac966eaf7a82baad25893105653f3)
- Update CHANGELOG.md and README.md file. [ba858d3](https://github.com/Sven-Seyfert/gomisc/commit/ba858d35f6ea20678d82592899175515da66a1ae)

## [0.2.0] - 2023-08-19

### Added

- Initialize go by go.mod file. [910ad1b](https://github.com/Sven-Seyfert/gomisc/commit/910ad1b3e3497eb293e91b54b34192cd800b20cd)
- Package 'singleinstance' and a usage example. [da79261](https://github.com/Sven-Seyfert/gomisc/commit/da792618f82908020a291555f17f9c56f5a1aa5f)

### Changed

- Extend .gitignore file. [efba4b4](https://github.com/Sven-Seyfert/gomisc/commit/efba4b41299360a67a58dc4dfe2c3c1783dc8051)

### Documented

- README.md structure change (first draft). [1b67221](https://github.com/Sven-Seyfert/gomisc/commit/1b67221e7d460ae6a9e434882011b4a6933ad2a9)
- Trivial upper case change in CHANGELOG.md for proper github username. [3d06708](https://github.com/Sven-Seyfert/gomisc/commit/3d0670837422b238a2a4aeb55fd088b8705d0425)

## [0.1.0] - 2023-08-19

### Added

- Initial commit [c151c43](https://github.com/Sven-Seyfert/gomisc/commit/c151c4395192ad655cef7355974abafb44f6ed33)

### Changed

- Delete unnecessary text in .gitignore file. [077b1e5](https://github.com/Sven-Seyfert/gomisc/commit/077b1e5e72eadc8a7c33f22450ade6b0f5bf74ed)

### Documented

- Add CHANGELOG.md file. [43a0db6](https://github.com/Sven-Seyfert/gomisc/commit/43a0db6463a0457cc45520d70dbbe9232aa2d622)
- Adjust style of LICENSE.md file. [bc6adc1](https://github.com/Sven-Seyfert/gomisc/commit/bc6adc1a8e8c52ad696c13c2c2b5167c2634d1a2)

[Unreleased]: https://github.com/sven-seyfert/gomisc/compare/v0.10.0...HEAD
[0.10.0]: https://github.com/sven-seyfert/gomisc/compare/v0.9.0...v0.10.0
[0.9.0]: https://github.com/sven-seyfert/gomisc/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/sven-seyfert/gomisc/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/sven-seyfert/gomisc/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/sven-seyfert/gomisc/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/sven-seyfert/gomisc/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/sven-seyfert/gomisc/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/sven-seyfert/gomisc/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/sven-seyfert/gomisc/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/sven-seyfert/gomisc/releases/tag/v0.1.0

---

### Legend - Types of changes

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Documented` for documentation only changes.
- `Fixed` for any bug fixes.
- `Refactored` for changes that neither fixes a bug nor adds a feature.
- `Removed` for now removed features.
- `Security` in case of vulnerabilities.
- `Styled` for changes like whitespaces, formatting, missing semicolons etc.

##

[To the top](#)
