---
name: "sdlc-ci-cd"
description: "执行软件开发生命周期中的集成阶段标准化规范。包含CI/CD流水线、版本控制、容器化与微服务部署等。当用户需要设计流水线、处理依赖或进行容器化时调用。"
---

# SDLC-CI-CD: 集成阶段标准化技能体系

本技能模块旨在为软件开发生命周期的【集成阶段】提供全方位的标准化规范和最佳实践指南。涵盖持续集成/持续交付（CI/CD）流水线设计、版本控制策略、依赖管理、环境配置、容器化部署、微服务集成等关键技术。

## 1. 多维度适配指南 (组合调用矩阵)

根据您的项目特性，应用以下推荐配置：

### 按项目类型适配
- **Web应用**: 强制实施前端资源打包（Webpack/Vite）、CDN同步、SourceMap分析。部署为SPA（Nginx）或SSR（Node服务）。
- **移动应用**: 搭建Mac Mini节点用于iOS打包，Android使用Gradle缓存。使用Fastlane进行应用商店分发，集成崩溃日志SDK。
- **大数据系统**: 针对Spark/Flink等任务，将JAR包上传至HDFS/S3，使用Airflow/DolphinScheduler进行工作流调度。
- **AI平台**: 针对模型训练使用Kubeflow/MLflow；Docker镜像需包含CUDA依赖（Nvidia-Docker）；模型权重文件存放在对象存储而非Git。

### 按技术栈适配
- **Java/Spring**: Maven/Gradle构建，生成Fat JAR或War包，利用Jib/Kaniko进行无守护进程镜像构建。
- **Python/Django**: 构建Wheel包，生成requirements.txt或Pipfile.lock，利用Alpine或Slim基础镜像减小体积。
- **Node.js**: 使用npm/yarn/pnpm生成lock文件，执行`npm ci`而非`npm install`，使用多阶段构建（Multi-stage Build）。
- **Go**: CGO_ENABLED=0，编译静态二进制文件，结合Scratch基础镜像，实现极小体积的容器化。
- **.NET**: 使用dotnet build/publish命令，利用微软官方.NET Runtime基础镜像，分层构建。

### 按团队规模适配
- **初创团队 (1-10人)**: 优先采用SaaS CI/CD服务（GitHub Actions, GitLab CI (Shared Runners), CircleCI），降低维护成本。
- **中型团队 (11-50人)**: 自建Jenkins集群或GitLab Runners，统一CI/CD模板（Pipeline as Code），使用Docker Swarm或基础K8s集群。
- **大型组织 (50人+)**: 企业级Kubernetes（Rancher/OpenShift），ArgoCD/Flux实施GitOps，Spinnaker进行多云/混合云发布管理。

### 按合规要求适配
- **等保 (2.0)**: CI/CD平台及节点需纳入主机安全防护；强制进行镜像安全扫描（Trivy/Clair），禁止容器使用特权模式（Privileged）。
- **GDPR**: 确保测试数据不包含生产PII，流水线日志脱敏（过滤Token/密码）。
- **SOX**: 职责分离，开发者无权直接触发生产环境发布。所有上线操作必须通过带数字签名的审批流进行审计。

---

## 2. 核心技能模块定义

### 2.1 持续集成/持续交付流水线设计 (CI/CD Pipeline)
*   **实施步骤**:
    1. Pipeline as Code: 将流水线配置（如`.gitlab-ci.yml`或`Jenkinsfile`）与源代码同仓管理。
    2. CI阶段: Checkout -> Lint -> Unit Test -> SonarQube Scan -> Build -> Build Image -> Push Registry。
    3. CD阶段: Deploy to Dev -> E2E Test -> Manual Approval -> Deploy to Staging -> Deploy to Prod。
*   **工具推荐**: GitLab CI/CD, GitHub Actions, Jenkins, Tekton, ArgoCD.
*   **质量门禁指标**: 构建时长<10分钟，单元测试覆盖率门禁（例如>80%），静态扫描无新增严重漏洞。
*   **风险控制措施**: 对敏感凭证（Secrets/Tokens）使用外部保险箱（HashiCorp Vault 或 AWS Secrets Manager）进行注入，禁止硬编码。
*   **交付物模板**: 《CI/CD流水线设计图》, 统一Pipeline模板库。

### 2.2 版本控制策略与分支管理 (Version Control Strategy)
*   **实施步骤**:
    1. 制定分支模型：GitFlow（适合定期发布），GitHub Flow（适合持续部署），Trunk-based（主干开发，需配合Feature Toggles）。
    2. 规范Commit Message（遵循Angular规范：feat, fix, docs, refactor, chore）。
    3. 语义化版本控制（Semantic Versioning, SemVer: MAJOR.MINOR.PATCH）。
*   **工具推荐**: Git, Commitizen, Standard Version, semantic-release.
*   **验收标准**: 所有的提交可追溯至工单（Jira ID/Issue Number），版本号变更触发Release Tag及Changelog自动生成。

### 2.3 依赖管理与制品库 (Dependency & Artifact Management)
*   **实施步骤**:
    1. 启用私有制品库，代理并缓存公共包管理工具（npm, maven, pypi）。
    2. 定期使用Dependabot/Renovate更新依赖。
    3. 锁定依赖版本（Lock files），避免幽灵依赖（Phantom Dependencies）。
*   **工具推荐**: Nexus, JFrog Artifactory, Harbor (Docker Registry).
*   **质量门禁指标**: 无已知高危CVE的依赖包，内部包引用规范通过率100%。

### 2.4 容器化部署与微服务集成 (Containerization & Microservices)
*   **实施步骤**:
    1. 编写Dockerfile，采用多阶段构建，使用非root用户运行容器（USER指令）。
    2. 配置K8s清单（Deployments, Services, ConfigMaps, Secrets）或使用Helm Charts。
    3. 定义微服务间的服务发现与负载均衡，集成服务网格（Service Mesh）。
*   **工具推荐**: Docker, Kubernetes, Helm, Kustomize, Istio/Linkerd.
*   **风险控制措施**: 配置资源限制（Requests/Limits），配置健康检查（Liveness/Readiness Probes），避免雪崩。
*   **交付物模板**: 基础架构即代码（IaC - Terraform/Ansible）配置文件，Helm Chart包。

### 2.5 环境配置与治理 (Environment Configuration)
*   **实施步骤**:
    1. 遵循12-Factor App原则，将配置与代码分离。
    2. 使用配置中心进行动态配置下发，区分多环境（Dev, Test, UAT, Prod）。
*   **工具推荐**: Apollo, Nacos, Spring Cloud Config, Consul.
*   **验收标准**: 环境切换无需重新编译代码（Build Once, Deploy Anywhere）。

---

## 3. 调用指南

当收到如下任务时，系统会自动应用本SKILL：
- “帮我写一个基于GitHub Actions的CI/CD流水线。”
- “这个Node项目如何进行Docker容器化，并提供K8s部署文件？”
- “如何规范我们团队的Git分支管理策略？”