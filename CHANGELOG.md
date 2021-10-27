CHANGELOG
=========

## HEAD (Unreleased)
* Upgrade to v3.2.0 of the Equinix Metal Terraform Provider.

---

## 3.0.0 (2021-09-27)
* Upgrade to v3.1.0 of the Equinix Metal Terraform Provider.
  **Please Note:** the following breaking change: 
  * `equinix-metal.getProject` now has no `bgpConfig` output

## 2.0.0 (2021-04-19)
* Depend on Pulumi 3.0, which includes improvements to Python resource arguments and key translation, Go SDK performance,
  Node SDK performance, general availability of Automation API, and more.
* Upgrade to v2.0.1 of the Equinix Metal Terraform Provider.
  **Please Note:**  
  This release introduces the new concept of "Equinix Metal Metros". You can read more about this in the [Equinix Metal documentation](https://feedback.equinixmetal.com/changelog/new-metros-feature-live)

## 1.4.0 (2021-04-12)
* Upgrade to pulumi-terraform-bridge v2.23.0

## 1.3.0 (2021-04-05)
* Upgrade to v1.1.0 of the Equinix Metal Terraform Provider

## 1.2.1 (2021-03-23)
* Upgrade to pulumi-terraform-bridge v2.22.1  
  **Please Note:** This includes a bug fix to the refresh operation regarding arrays

## 1.2.0 (2021-03-15)
* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 1.1.2 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider
* Avoid storing config from the environment into the state

## 1.1.1 (2021-02-09)
* Regenerate Go SDK to fix examples generation

## 1.1.0 (2021-01-29)
* Upgrade to pulumi-terraform-bridge v2.18.1

## 1.0.1 (2021-01-13)
* Update schema for facilities, plans, and operating systems.
* Upgrade to pulumi-terraform-bridge v2.17.0
* Upgrade to Pulumi v2.17.0

## 1.0.0 (2020-12-23)
* Initial creation of the provider against v1.0.0 of the Equinix Metal Terraform Provider.
