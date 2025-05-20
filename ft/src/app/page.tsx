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

// function ServerDataDisplay() {
//   const [serverData, setServerData] = useState<string | null>(null);

//   useEffect(() => {
//     fetch("http://localhost:8080/", {
//       method: "GET",
//       headers: {
//         "Content-Type": "text/plain",
//       },
//     })
//       .then((response) => {
//         if (!response.ok) {
//           throw new Error("Failed to fetch server data");
//         }
//         return response.text();
//       })
//       .then((data) => setServerData(data))
//       .catch((error) => console.error("Error fetching server data:", error));
//   }, []);

//   return (
//     <div>
//       <h3>服务器数据展示</h3>
//       <pre>{serverData || "加载中..."}</pre>
//     </div>
//   );
// }
function KeyDisplay() {
  const [keyData, setKeyData] = useState<string | null>(null);

  useEffect(() => {
    fetch("http://localhost:8080/key", {
      method: "GET",
      headers: {
        "Content-Type": "text/plain",
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to fetch key data");
        }
        return response.text();
      })
      .then((data) => setKeyData(data))
      .catch((error) => console.error("Error fetching key data:", error));
  }, []);

  return (
    <div>
      <pre>{keyData || "加载中..."}</pre>
    </div>
  );
}

function CsrDisplay() {
  const [csrData, setCsrData] = useState<string | null>(null);

  useEffect(() => {
    fetch("http://localhost:8080/csr", {
      method: "GET",
      headers: {
        "Content-Type": "text/plain",
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to fetch CSR");
        }
        return response.text();
      })
      .then((data) => setCsrData(data))
      .catch((error) => console.error("Error fetching CSR:", error));
  }, []);

  return (
    <div>
      <h3>客户端生成CSR并发给RA验证</h3>
      <pre>{csrData || "加载中..."}</pre>
    </div>
  );
}

function CertDisplay() {
  const [certData, setCertData] = useState<string | null>(null);

  useEffect(() => {
    fetch("http://localhost:8080/cert", {
      method: "GET",
      headers: {
        "Content-Type": "text/plain",
      },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Failed to fetch certificate");
        }
        return response.text();
      })
      .then((data) => setCertData(data))
      .catch((error) => console.error("Error fetching certificate:", error));
  }, []);

  return (
    <div>
      <h3>通过验证的CSR会发给CA，最后签发证书</h3>
      <pre>{certData || "加载中..."}</pre>
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
      <pre>{caKey || "加载中..."}</pre>
    </div>
  );
}

function ReturnType() {
  return (
    <div style={{ padding: "20px", fontFamily: "Arial, sans-serif" }}>
      <h1 style={{ textAlign: "center", marginBottom: "20px" }}>系统信息展示</h1>
      <div style={{ display: "flex", gap: "20px", flexWrap: "wrap" }}>
        <div style={{ flex: "1", border: "1px solid #ccc", padding: "10px", borderRadius: "8px", backgroundColor: "#f9f9f9" }}>
          <h2>客户端</h2>
          <KeyDisplay />
          <CsrDisplay />
          <CertDisplay />
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

function ReturnType1() {
  return (
    <div style={{ padding: "20px", fontFamily: "Arial, sans-serif" }}>
      {(() => {
        const steps = [
          <div key="server-list" className="flex-1 border border-gray-300 p-2.5 rounded-lg bg-gray-100">
            <ServerList />
          </div>,
          <div key="ca-key" className="flex-1 border border-gray-300 p-2.5 rounded-lg bg-gray-100">
            <CaKeyDisplay />
          </div>,
          <div key="ca-cert" className="flex-1 border border-gray-300 p-2.5 rounded-lg bg-gray-100">
            <CaCertDisplay />
          </div>,
          <div key="key" className="flex-1 border border-gray-300 p-2.5 rounded-lg bg-gray-100">
            <KeyDisplay />
          </div>,
          <div key="csr" className="flex-1 border border-gray-300 p-2.5 rounded-lg bg-gray-100">
            <CsrDisplay />
          </div>,
          <div key="cert" className="flex-1 border border-gray-300 p-2.5 rounded-lg bg-gray-100">
            <CertDisplay />
          </div>,
        ];
        const [step, setStep] = useState(0);
        return (
          <div>
            <div style={{ display: "flex", gap: "20px", flexWrap: "wrap" }}>
              {steps.slice(0, step + 1)}
            </div>
            {step < steps.length - 1 && (
              <button className=" bg-blue-400 rounded-lg"
                style={{ marginTop: 20, padding: "8px 16px", fontSize: 16 }}
                onClick={() => setStep(step + 1)}
              >
                Next
              </button>
            )}
          </div>
        );
      })()}
    </div>
  );
}

function ReturnType2() {
  return (
    <div>
      <div>
        <div className="border border-blue-300 p-4 rounded-4xl bg-blue-400">初始化工作区</div>
        <div className="p-4">
          <details className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">
            <summary>部署服务注册中心</summary>
            <a
              href="http://localhost:3001/servers"
              target="_blank"
              rel="noopener noreferrer"
              style={{ color: "#2563eb", textDecoration: "underline", marginLeft: 8 }}
            >
              链接
            </a>
            <ServerList />
          </details>
          <details className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">
            <summary>部署CA</summary>
            <details className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">
              <summary>CA私钥</summary>
              <CaKeyDisplay />
            </details>
            <details className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">
              <summary>CA证书</summary>
              <CaCertDisplay />
            </details>
          </details>
          {/* <details className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">
            <summary>部署RA</summary>
            <ServerList />
          </details> */}
          <div className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">部署RA</div>
        </div>
      </div>
      <div>
        <div className="border border-blue-300 p-4 rounded-4xl bg-blue-400">工作流程开始</div>
        <div style={{ padding: "20px", fontFamily: "Arial, sans-serif" }}>
          {(() => {
            const steps = [
              <div key="key" className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">
                <details className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">
                  <summary>客户端本地生成密钥对</summary>
                  <KeyDisplay />
                </details>
              </div>,
              <div key="csr" className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">
                <details className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">
                  <summary>客户端本地CSR</summary>
                  <CsrDisplay />
                </details>
              </div>,
              <div key="ra" className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">
                <div className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">
                  RA验证CSR通过后转发CA
                </div>
              </div>,
              <div key="cert" className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">
                <details className="border border-gray-300 p-2.5 rounded-lg bg-gray-100">
                  <summary>CA签发证书并发送给客户端</summary>
                  <CertDisplay />
                </details>
              </div>,
            ];
            const [step, setStep] = useState(0);
            return (
              <div>
                <div className=" gap-[20px]"
                // style={{ display: "flex", gap: "20px", flexWrap: "wrap" }}
                >
                  {steps.slice(0, step + 1)}
                </div>
                {step < steps.length - 1 && (
                  <button className=" bg-blue-400 rounded-lg"
                    style={{ marginTop: 20, padding: "8px 16px", fontSize: 16 }}
                    onClick={() => setStep(step + 1)}
                  >
                    Next
                  </button>
                )}
              </div>
            );
          })()}
        </div>
      </div>
    </div>
  );
}
export default function Home() {

  return (
    // <ReturnType />
    // <ReturnType1 />
    <ReturnType2 />
  );
}
