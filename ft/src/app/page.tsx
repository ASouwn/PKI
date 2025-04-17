"use client";
import { useEffect, useState } from "react";

// 服务注册中心数据展示
function ServerList() {
  const [servers, setServers] = useState<Record<string, any>>({});

  useEffect(() => {
    fetch("http://localhost:3001/servers", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .catch((error) => {
        console.error("Network error or CORS issue:", error);
        throw error;
      })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to fetch server data");
        }
        return response.json();
      })
      .then((data) => setServers(data))
      .catch((error) => console.error("Error fetching server data:", error));
  }, []);

  return (
    <div>
      <h3>服务器列表</h3>
      <ul>
        {Object.entries(servers).map(([key, value]) => (
          <li key={key}>
            <strong>{key}</strong>: {JSON.stringify(value)}
          </li>
        ))}
      </ul>
    </div>
  );
}

function ServerDataDisplay() {
  const [serverData, setServerData] = useState<string | null>(null);

  useEffect(() => {
    fetch("http://localhost:8080/", {
      method: "GET",
      headers: {
        "Content-Type": "text/plain",
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to fetch server data");
        }
        return response.text();
      })
      .then((data) => setServerData(data))
      .catch((error) => console.error("Error fetching server data:", error));
  }, []);

  return (
    <div>
      <h3>服务器数据展示</h3>
      <pre>{serverData || "加载中..."}</pre>
    </div>
  );
}

function CaCertDisplay() {
  const [caCert, setCaCert] = useState<string | null>(null);

  useEffect(() => {
    fetch("http://localhost:8082/ca-cert", {
      method: "GET",
      headers: {
        "Content-Type": "text/plain",
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to fetch CA certificate");
        }
        return response.text();
      })
      .then((data) => setCaCert(data))
      .catch((error) => console.error("Error fetching CA certificate:", error));
  }, []);

  return (
    <div>
      <h3>CA 证书</h3>
      <pre>{caCert || "加载中..."}</pre>
    </div>
  );
}

function CaKeyDisplay() {
  const [caKey, setCaKey] = useState<string | null>(null);

  useEffect(() => {
    fetch("http://localhost:8082/ca-key", {
      method: "GET",
      headers: {
        "Content-Type": "text/plain",
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to fetch CA private key");
        }
        return response.text();
      })
      .then((data) => setCaKey(data))
      .catch((error) => console.error("Error fetching CA private key:", error));
  }, []);

  return (
    <div>
      <h3>CA 私钥</h3>
      <pre>{caKey || "加载中..."}</pre>
    </div>
  );
}

export default function Home() {
  return (
    <div style={{ padding: "20px", fontFamily: "Arial, sans-serif" }}>
      <h1 style={{ textAlign: "center", marginBottom: "20px" }}>系统信息展示</h1>
      <div style={{ display: "flex", gap: "20px", flexWrap: "wrap" }}>
        <div style={{ flex: "1", border: "1px solid #ccc", padding: "10px", borderRadius: "8px", backgroundColor: "#f9f9f9" }}>
          <h2>客户端</h2>
          <ServerDataDisplay />
        </div>
        <div style={{ flex: "1", border: "1px solid #ccc", padding: "10px", borderRadius: "8px", backgroundColor: "#f9f9f9" }}>
          <h2>CA</h2>
          <CaCertDisplay />
          <CaKeyDisplay />
        </div>
        <div style={{ flex: "1", border: "1px solid #ccc", padding: "10px", borderRadius: "8px", backgroundColor: "#f9f9f9" }}>
          <h2>服务注册中心</h2>
          <ServerList />
        </div>
      </div>
    </div>
  );
}
