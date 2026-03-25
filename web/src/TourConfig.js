import React from "react";

export const TourObj = {
  home: [
    {
      title: "Welcome to CheersAI-SSO",
      description: "You can use CheersAI-SSO to manage sign-in, organization, and identity workflows.",
      cover: (
        <img
          alt="CheersAI-SSO"
          src={"/logo.png"}
        />
      ),
    },
    {
      title: "Statistic cards",
      description: "Here are four statistic cards for user information.",
      id: "statistic",
    },
    {
      title: "Import users",
      description: "You can add new users or update existing users by uploading a XLSX file of user information.",
      id: "echarts-chart",
    },
  ],
  webhooks: [
    {
      title: "Webhook List",
      description: "Event systems allow you to build integrations that subscribe to signup, login, logout, and user update events. When an event is triggered, the system sends a POST json payload to the configured URL.",
    },
  ],
  syncers: [
    {
      title: "Syncer List",
      description: "CheersAI-SSO stores users in the user table and provides syncers to help migrate external user data quickly.",
    },
  ],
  sysinfo: [
    {
      title: "CPU Usage",
      description: "You can see the CPU usage in real time.",
      id: "cpu-card",
    },
    {
      title: "Memory Usage",
      description: "You can see the Memory usage in real time.",
      id: "memory-card",
    },
    {
      title: "API Latency",
      description: "You can see the usage statistics of each API latency in real time.",
      id: "latency-card",
    },
    {
      title: "API Throughput",
      description: "You can see the usage statistics of each API throughput in real time.",
      id: "throughput-card",
    },
    {
      title: "About CheersAI-SSO",
      description: "You can get more product information in this card.",
      id: "about-card",
    },
  ],
  subscriptions: [
    {
      title: "Subscription List",
      description: "Subscription helps to manage user's selected plan that make easy to control application's features access.",
    },
  ],
  pricings: [
    {
      title: "Price List",
      description: "CheersAI-SSO can be used as a subscription management system via plan, pricing and subscription.",
    },
  ],
  plans: [
    {
      title: "Plan List",
      description: "Plan describes a list of application features with its own name and price. Plan features depend on roles and permissions, and can vary by region or date.",
    },
  ],
  payments: [
    {
      title: "Payment List",
      description: "After the payment is successful, you can see the transaction information of the products in Payment, such as organization, user, purchase time, product name, etc.",
    },
  ],
  products: [
    {
      title: "Session List",
      description: "You can add the product (or service) you want to sell. The following will tell you how to add a product.",
    },
  ],
  sessions: [
    {
      title: "Session List",
      description: "You can get Session ID in this list.",
    },
  ],
  tokens: [
    {
      title: "Token List",
      description: "CheersAI-SSO is based on OAuth. Tokens are users' OAuth tokens, and you can view access tokens in this list.",
    },
  ],
  enforcers: [
    {
      title: "Enforcer List",
      description: "In addition to the API interface for permission enforcement, CheersAI-SSO also provides interfaces that help external applications obtain permission policy information.",
    },
  ],
  adapters: [
    {
      title: "Adapter List",
      description: "The UI supports connecting adapters and managing policy rules. Adapters can be used to load policy rules from storage or save policy rules back to it.",
    },
  ],
  models: [
    {
      title: "Model List",
      description: "Model defines your permission policy structure, and how requests should match these permission policies and their effects. Then you can user model in Permission.",
    },
  ],
  permissions: [
    {
      title: "Permission List",
      description: "All users associated with a single organization are shared between the organization's applications and therefore have access to the applications. You can use permissions to restrict access to certain applications or resources.",
    },
    {
      title: "Permission Add",
      description: "In the web UI, you can add a Model for your organization in the Model configuration item, and a Policy for your organization in the Permission configuration item.",
      id: "add-button",
    },
    {
      title: "Permission Upload",
      description: "You can import model and policy files that match your usage scenario through the web UI.",
      id: "upload-button",
    },
  ],
  roles: [
    {
      title: "Role List",
      description: "Each user may have multiple roles. You can see the user's roles on the user's profile.",
    },
  ],
  resources: [
    {
      title: "Resource List",
      description: "You can upload resources in CheersAI-SSO. Before uploading resources, you need to configure a storage provider.",
    },
    {
      title: "Upload Resource",
      description: "Users can upload resources such as files and images to the previously configured cloud storage.",
      id: "upload-button",
    },
  ],
  providers: [
    {
      title: "Provider List",
      description: "We have 6 kinds of providers:OAuth providers、SMS Providers、Email Providers、Storage Providers、Payment Provider、Captcha Provider.",
    },
    {
      title: "Provider Add",
      description: "You must add the provider to application, then you can use the provider in your application",
      id: "add-button",
    },
  ],
  organizations: [
    {
      title: "Organization List",
      description: "Organization is the basic unit of CheersAI-SSO, which manages users and applications. If a user signs in to an organization, the user can access its applications without signing in again.",
    },
  ],
  groups: [
    {
      title: "Group List",
      description: "In the groups list pages, you can see all the groups in organizations.",
    },
  ],
  users: [
    {
      title: "User List",
      description: "As an authentication platform, CheersAI-SSO is able to manage users.",
    },
    {
      title: "Import users",
      description: "You can add new users or update existing users by uploading a XLSX file of user information.",
      id: "upload-button",
    },
  ],
  applications: [
    {
      title: "Application List",
      description: "If you want to use CheersAI-SSO to provide login service for your web apps, you can add them as applications. Users can access all applications in their organizations without signing in twice.",
    },
  ],
};

export const TourUrlList = ["home", "organizations", "groups", "users", "applications", "providers", "resources", "roles", "permissions", "models", "adapters", "enforcers", "tokens", "sessions", "products", "payments", "plans", "pricings", "subscriptions", "sysinfo", "syncers", "webhooks"];

export function getNextUrl(pathName = window.location.pathname) {
  return TourUrlList[TourUrlList.indexOf(pathName.replace("/", "")) + 1] || "";
}

let orgIsTourVisible = true;

export function setOrgIsTourVisible(visible) {
  orgIsTourVisible = visible;
  if (orgIsTourVisible === false) {
    setIsTourVisible(false);
  }
}

export function setIsTourVisible(visible) {
  localStorage.setItem("isTourVisible", visible);
  window.dispatchEvent(new Event("storageTourChanged"));
}

export function setTourLogo(tourLogoSrc) {
  if (tourLogoSrc !== "") {
    TourObj["home"][0]["cover"] = (<img alt="casdoor.png" src={tourLogoSrc} />);
  }
}

export function getTourVisible() {
  return localStorage.getItem("isTourVisible") !== "false";
}

export function getNextButtonChild(nextPathName) {
  return nextPathName !== "" ?
    `Go to "${nextPathName.charAt(0).toUpperCase()}${nextPathName.slice(1)} List"`
    : "Finish";
}

export function getSteps() {
  const path = window.location.pathname.replace("/", "");
  const res = TourObj[path];
  if (res === undefined) {
    return [];
  } else {
    return res;
  }
}
