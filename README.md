<h1 align="center" style="border-bottom: none;">📦⚡️ CheersAI-SSO</h1>
<h3 align="center">CheersAI 全平台统一身份认证（SSO）与访问控制中心</h3>

<p align="center">
  <em>基于先进的 OAuth2 / OIDC 协议，为 CheersAI 办公产品矩阵（Desktop、Vault、Nexus 等）提供统一、安全、高效的身份认证服务。</em>
</p>

## 📖 项目概述

**CheersAI-SSO** 是 [CheersAI](https://cheersai.cloud) 专用的集中式身份管理平台。本系统在优秀的开源项目 Casdoor 基础上进行了深度的品牌定制与 UI 重构，以完美契合 CheersAI 全线产品的交互规范与安全需求。

### ✨ 核心特性

- **一处登录，全家通行 (SSO)**：连接 CheersAI Desktop、Nexus 及其他子系统，用户只需进行一次认证。
- **CheersAI 专属视觉**：采用 CheersAI 官方 UI/UE 规范的主题色、组件和布局体验，从登录页到管理后台的无缝品牌融合。
- **现代化多协议支持**：全面支持 OAuth 2.1, OIDC, SAML 等业界主流协议，保障多端（Web / Tauri / App）的顺畅接入。
- **高安全标准**：包含完善的审计日志、多因素认证（MFA）、会话防泄漏等机制。

---

## 🚀 如何接入 (For Developers)

无论你负责的是 Next.js 的前端项目（如 Nexus），还是基于 Tauri / Rust 开发的客户端（如 Desktop），都可以通过轻量级 SDK 快速接入。

> 详细的系统接入指南、API 调用及 Token 刷新机制，请查阅完整文档：
> [👉 点击这里查看《CheersAI-SSO系统接入说明文档》](./docs/CheersAI-SSO系统接入说明文档.md)

### 快速开始 (Web / Node.js 示例)
```bash
npm install casdoor-js-sdk
```

---

## 🛠️ 环境部署合规

* **生产环境端点**: `https://sso.cheersai.cloud`
* **组织标识 (Organization)**: `CheersAI`

*注：本系统基于 Go 语言开发，支持 Docker 镜像一键私有化部署。如需本地开发调试，请确保正确配置了 `.env` 或 `conf/app.conf` 中的数据库连接*。

---

## 📝 鸣谢与许可

本项目底层核心代码基于优秀的开源项目 [Casdoor](https://github.com/casdoor/casdoor)。在遵循其开源精神的前提下，CheersAI 团队进行了产品化与品牌化的重构。

- 核心引擎协议: [Apache-2.0](https://github.com/casdoor/casdoor/blob/master/LICENSE)
- 版权所有 © 2026 CheersAI Team.
