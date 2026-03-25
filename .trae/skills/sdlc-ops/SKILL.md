---
name: "sdlc-ops"
description: "执行软件开发生命周期中的发布运维阶段标准化规范。包含发布策略、监控告警、日志管理、容量规划与灾备。当需要设计高可用架构、监控方案或发布回滚机制时调用。"
---

# SDLC-Ops: 发布运维阶段标准化技能体系

本技能模块旨在为软件开发生命周期的【发布运维阶段】提供全方位的标准化规范和最佳实践指南。涵盖发布策略制定、灰度发布、回滚机制、监控告警、日志管理、容量规划、灾备方案等运维技能。

## 1. 多维度适配指南 (组合调用矩阵)

根据您的项目特性，应用以下推荐配置：

### 按项目类型适配
- **Web应用**: 蓝绿部署为主，重点监控HTTP状态码（5xx/4xx）、页面首屏时间（LCP）、前端JS错误率。
- **移动应用**: 采用分阶段发布（Phased Release），结合热修复（Hotfix/JSPatch），监控Crash率、ANR和API成功率。
- **大数据系统**: 离线任务调度监控（Airflow DAG失败告警），存储水位监控（HDFS/S3），计算资源倾斜（Data Skew）告警。
- **AI平台**: A/B测试为主（验证模型效果），重点监控推理延迟（Latency P99）、GPU利用率、显存溢出（OOM）及模型漂移（Data Drift）。

### 按技术栈适配
- **Java/Spring**: 接入Prometheus + Micrometer，通过Grafana展示JVM GC停顿、线程池打满、JDBC连接池状态。
- **Python/Django**: 监控Gunicorn/uWSGI worker状态，使用Sentry捕获未处理异常，追踪Celery任务队列积压。
- **Node.js**: 监控Event Loop Lag，V8内存使用量，使用PM2/Clinic.js进行进程级监控与守护。
- **Go**: 暴露`/debug/pprof`（仅限内网），监控Goroutine数量，利用OpenTelemetry进行全链路追踪。
- **.NET**: 收集EventCounters，结合Application Insights，监控Kestrel连接数和GC指标。

### 按团队规模适配
- **初创团队 (1-10人)**: 停机发布或简单的滚动更新（Rolling Update），依赖云厂商原生监控（AWS CloudWatch/阿里云监控），日志直接写入本地或云托管服务。
- **中型团队 (11-50人)**: 金丝雀发布（Canary Release），自建ELK/EFK日志堆栈，Prometheus+Grafana+Alertmanager告警体系，定义On-call轮值。
- **大型组织 (50人+)**: 全链路流量染色，异地多活，基于AIOps的异常检测，混沌工程常态化，严格的SLI/SLO/SLA指标管理体系。

### 按合规要求适配
- **等保 (2.0)**: 日志必须留存6个月以上，严禁篡改（WORM存储），发布操作需多因素认证（MFA）与堡垒机审计。
- **GDPR**: 运维排障时禁止导出包含PII的生产数据至本地，日志系统需自动掩码（Masking）手机号、邮箱等敏感信息。
- **SOX**: 所有线上变更（发布/回滚/配置修改）必须有审批工单（Change Request），不可由单一工程师闭环完成。

---

## 2. 核心技能模块定义

### 2.1 发布策略制定与执行 (Release Strategies)
*   **实施步骤**:
    1. 评估发布影响面，选择发布模型：滚动发布（Rolling）、蓝绿发布（Blue/Green）、金丝雀/灰度发布（Canary）。
    2. 配置流量切分规则（如基于Header、地域、用户UID的Nginx/Istio路由规则）。
    3. 执行数据库迁移（Schema Migration），采用向前兼容设计（Expand-Contract Pattern），避免锁表。
*   **工具推荐**: Argo Rollouts, Spinnaker, Istio, Flagger, Flyway/Liquibase.
*   **质量门禁指标**: 灰度阶段错误率增量<0.1%，无严重线上客诉。
*   **交付物模板**: 《发布操作指导书》(Runbook)，变更申请单（RFC）。
*   **验收标准**: 新旧版本平滑过渡，用户无感知断线（Zero Downtime Deployment）。

### 2.2 回滚机制与故障恢复 (Rollback & Recovery)
*   **实施步骤**:
    1. 制定“一键回滚”策略，确保容器镜像Tag与Git Commit绑定。
    2. 对于数据库变更，准备回滚脚本（Down Script），或依赖数据快照（Snapshot）/增量备份恢复。
    3. 演练MTTR（平均恢复时间）。
*   **工具推荐**: K8s Rollout Undo, ArgoCD Sync, RDS快照恢复.
*   **风险控制措施**: 禁止执行不可逆的数据删除（Drop/Truncate），采用软删除（Soft Delete）。

### 2.3 监控告警体系 (Monitoring & Alerting)
*   **实施步骤**:
    1. 定义SLI（服务质量指标），如可用性、延迟、吞吐量。设定SLO（服务质量目标），如99.9%。
    2. 实施USE方法论（使用率、饱和度、错误）监控基础设施，RED方法论（请求速率、错误、持续时间）监控服务。
    3. 分级告警：P0（电话/短信），P1（IM群通知），P2（邮件）。
*   **工具推荐**: Prometheus, Grafana, Alertmanager, PagerDuty, Zabbix.
*   **质量门禁指标**: 告警准确率>90%（减少狼来了效应），MTTD（平均发现时间）< 5分钟。
*   **交付物模板**: 《监控指标仪表盘设计》，On-call轮值表与升级策略（Escalation Policy）。

### 2.4 日志管理与全链路追踪 (Logging & Tracing)
*   **实施步骤**:
    1. 统一日志格式（JSON），注入TraceID/SpanID。
    2. 集中式采集（Filebeat/Fluentd），解析过滤（Logstash/Vector），存储检索（Elasticsearch/Loki）。
    3. 接入APM（应用性能管理）追踪跨微服务调用链路。
*   **工具推荐**: ELK/EFK Stack, Loki, SkyWalking, Jaeger, OpenTelemetry.
*   **验收标准**: 可通过TraceID在日志平台中检索出整个跨服务请求的完整生命周期日志。

### 2.5 容量规划与灾备方案 (Capacity Planning & Disaster Recovery)
*   **实施步骤**:
    1. 容量压测（基于历史峰值x倍数），配置弹性伸缩（HPA/VPA/Karpenter）。
    2. 制定RTO（恢复时间目标）和RPO（恢复点目标）。
    3. 设计同城双活/异地多活，定期进行容灾演练（切换DNS或流量网关）。
*   **工具推荐**: K8s HPA, AWS Auto Scaling, Chaos Mesh, Sentinel (限流降级).
*   **风险控制措施**: 关键数据异地冷备，实施降级开关（Feature Toggles），在极端情况下保核心链路。
*   **交付物模板**: 《容灾演练报告》，《容量规划与成本预估表》。

---

## 3. 调用指南

当收到如下任务时，系统会自动应用本SKILL：
- “帮我设计一个高并发Web应用的监控和告警体系。”
- “如何实现K8s上的平滑蓝绿发布和一键回滚？”
- “系统准备双十一大促，需要做哪些容量规划和降级预案？”