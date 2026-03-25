---
name: "sdlc-test"
description: "执行软件开发生命周期中的测试阶段标准化规范。包含单元测试、集成测试、自动化框架搭建与性能安全测试等。当需要编写测试用例、自动化测试或缺陷管理时调用。"
---

# SDLC-Test: 测试阶段标准化技能体系

本技能模块旨在为软件开发生命周期的【测试阶段】提供全方位的标准化规范和最佳实践指南。涵盖单元测试、集成测试、系统测试、自动化测试框架、缺陷管理、性能与安全测试等核心环节。

## 1. 多维度适配指南 (组合调用矩阵)

根据您的项目特性，应用以下推荐配置：

### 按项目类型适配
- **Web应用**: 强制实施E2E测试（Selenium/Playwright/Cypress），跨浏览器兼容性测试，首屏加载性能测试。
- **移动应用**: 强制真机/模拟器矩阵测试（Appium），弱网测试（Charles/Fiddler），耗电量及发热测试。
- **大数据系统**: 数据质量断言（Great Expectations），大规模并发查询压测，批流数据一致性验证。
- **AI平台**: 模型鲁棒性测试，对抗样本注入测试，数据集偏见/偏差检测，推理延迟压测。

### 按技术栈适配
- **Java/Spring**: 单元测试使用JUnit 5 + Mockito；API集成测试使用REST Assured；性能测试JMeter。
- **Python/Django**: 单元测试使用Pytest；Web UI使用Playwright-Python；性能使用Locust。
- **Node.js**: 单元测试使用Jest/Vitest；端到端使用Cypress/Puppeteer；压测使用Artillery/k6。
- **Go**: 单元测试原生`testing`包 + `testify`；API测试`httpexpect`。
- **.NET**: 单元测试xUnit/NUnit + Moq；性能测试BenchmarkDotNet。

### 按团队规模适配
- **初创团队 (1-10人)**: 开发兼顾测试（TDD推荐），注重核心路径（Happy Path）的单元测试和轻量级接口自动化。
- **中型团队 (11-50人)**: 专职QA团队介入，搭建基于Jenkins/GitLab的自动化测试流水线，实施每日构建回归。
- **大型组织 (50人+)**: 测试左移与右移，混沌工程（Chaos Engineering）演练，全链路压测体系，AI辅助生成测试用例。

### 按合规要求适配
- **等保 (2.0)**: 强制渗透测试（黑盒测试），定期进行漏洞扫描（DAST/SAST/IAST）。
- **GDPR**: 确保测试环境禁止使用生产真实PII数据，强制执行数据脱敏后落库测试环境。
- **SOX**: 测试结果防篡改，强制缺陷生命周期完整追溯，审计所有UAT验收签名报告。

---

## 2. 核心技能模块定义

### 2.1 单元测试与测试驱动开发 (Unit Testing & TDD)
*   **实施步骤**:
    1. 遵循红-绿-重构（Red-Green-Refactor）节奏或后置补齐单元测试。
    2. 采用AAA模式（Arrange-Act-Assert）或Given-When-Then结构。
    3. Mock/Stub外部依赖，确保测试独立性和可重复执行性。
*   **工具推荐**: JUnit, Pytest, Jest, xUnit, Mockito, Sinon.
*   **质量门禁指标**: 核心业务逻辑行覆盖率(Line Coverage) > 80%，分支覆盖率(Branch Coverage) > 70%，测试执行时间<5分钟。
*   **风险控制措施**: 警惕无效的断言（Assert-Free Tests），避免测试代码与实现细节过度耦合（脆性测试）。
*   **交付物模板**: 《单元测试覆盖率报告》(JaCoCo, Istanbul, Cobertura)。
*   **验收标准**: 单元测试全数通过，无Flaky Tests（间歇性失败的测试）。

### 2.2 集成测试与系统测试 (Integration & System Testing)
*   **实施步骤**:
    1. 验证模块间接口交互（API契约测试）。
    2. 搭建真实的数据库/消息队列依赖（使用Testcontainers）。
    3. 执行跨服务业务链路流转验证。
*   **工具推荐**: Postman/Newman, REST Assured, Testcontainers, Pact (契约测试).
*   **风险控制措施**: 避免依赖不稳定外部第三方API，使用WireMock等工具进行隔离。
*   **交付物模板**: 《API集成测试报告》。

### 2.3 自动化测试框架搭建 (Automation Framework)
*   **实施步骤**:
    1. 采用分层测试金字塔策略（Unit > API > UI）。
    2. 实施数据驱动测试（DDT）和关键字驱动测试（KDT）。
    3. 剥离测试数据与测试脚本，统一配置管理环境（Dev/Test/Staging）。
*   **工具推荐**: Selenium, Playwright, Cypress, Appium, Robot Framework.
*   **质量门禁指标**: 回归测试自动化率>70%，核心主干流程100%自动化。

### 2.4 性能测试与安全测试 (Performance & Security Testing)
*   **实施步骤**:
    1. 性能：确定TPS/QPS、并发数、响应时间基线，进行负载测试、压力测试、容量测试。
    2. 安全：进行OWASP Top 10漏洞扫描，实施模糊测试（Fuzzing），权限越权测试。
*   **工具推荐**: JMeter, k6, Locust, Gatling; ZAP, Burp Suite, SonarQube (SAST).
*   **质量门禁指标**: P99响应时间<200ms（或依SLA定），无高危及以上安全漏洞（0 Critical/High CVEs）。
*   **验收标准**: 《性能压测报告》（含瓶颈分析），《渗透测试与漏洞扫描报告》。

### 2.5 缺陷管理与测试用例设计 (Defect Management & Test Case Design)
*   **实施步骤**:
    1. 使用等价类划分、边界值分析、状态迁移等黑盒方法设计用例。
    2. 规范Bug生命周期（New -> Assigned -> Fixed -> Verified -> Closed）。
*   **工具推荐**: Jira, TestRail, Xray, Bugzilla.
*   **交付物模板**: 《测试用例全集》(Test Suite)，《缺陷管理规范》。

---

## 3. 调用指南

当收到如下任务时，系统会自动应用本SKILL：
- “帮我为这个类生成全面的单元测试，包括边界条件。”
- “如何搭建一个基于Playwright的E2E测试框架？”
- “帮我设计一个接口的性能压测方案。”