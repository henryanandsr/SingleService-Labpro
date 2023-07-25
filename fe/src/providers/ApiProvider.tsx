import OHLBEAPI, { api as OHLBEAxios } from "@/api/ohl-be";
import { LoadingSpinner } from "@/components";
import { APIContext } from "@/contexts/";
import { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";

interface APIProviderProps {
  children: React.ReactNode;
}

export default function APIProvider({ children }: APIProviderProps) {
  const location = useLocation();
  const navigate = useNavigate();
  const [APIUrl, setURL] = useState<string | null>(null);
  if (APIUrl) {
    OHLBEAxios.defaults.baseURL = APIUrl;
  }

  const setToken = (token: string | null) => {
    if (token) {
      OHLBEAxios.defaults.headers.common["Authorization"] = `${token}`;
    } else {
      delete OHLBEAxios.defaults.headers.common["Authorization"];
    }
  };

  const setAPIUrl = (url: string) => {
    setURL(url);
    localStorage.setItem("api", url);
    OHLBEAxios.defaults.baseURL = url;
  };

  useEffect(() => {
    const url = localStorage.getItem("api");
    if (url) {
      setAPIUrl(url);
    }
    if (location.pathname !== "/config" && !url) {
      navigate("/config");
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [APIUrl, location.pathname]);

  if (location.pathname !== "/config" && !APIUrl) {
    return <LoadingSpinner />;
  }

  return (
    <APIContext.Provider
      value={{
        api: OHLBEAPI,
        APIUrl: APIUrl,
        setAPIUrl,
        token: OHLBEAxios.defaults.headers.common["Authorization"]?.toString() || null,
        setToken,
      }}
    >
      {children}
    </APIContext.Provider>
  );
}
