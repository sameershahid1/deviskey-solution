import { useEffect, useState } from "react";

function App() {
  const [downloading, setDownloading] = useState(false);
  const [error, setError] = useState<any>(null);

  useEffect(() => {
    (async () => {
      const raw = await fetch("http://localhost:8080/test");
      const response = await raw.json();
      console.log(response);
    })();
  }, []);

  const handleDownload = async () => {
    setDownloading(true);
    try {
      const response = await fetch("http://localhost:8080/generate-pdf"); // Adjust the endpoint path
      response.blob().then((blob) => {
        let url = window.URL.createObjectURL(blob);
        let a = document.createElement("a");
        a.href = url;
        a.download = "employees.pdf";
        a.click();
      });

      if (!response.ok) {
        throw new Error(`Failed to download PDF: ${response.statusText}`);
      }
    } catch (error: any) {
      setError(error);
    } finally {
      setDownloading(false);
    }
  };
  return (
    <div>
      <button onClick={handleDownload} disabled={downloading}>
        {downloading ? "Downloading..." : "Download PDF"}
      </button>
      {error && <p className="error">{error.message}</p>}
    </div>
  );
}

export default App;
