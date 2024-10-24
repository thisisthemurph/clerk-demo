import {useAuth} from "@clerk/clerk-react"

export default function useFetch() {
  const { getToken } = useAuth()

  return async (input: RequestInfo, init?: RequestInit) => {
    try {
      const token = await getToken();

      const headers = new Headers(init?.headers || {});
      if (token) {
        headers.append("Authorization", token);
      }

      const response = await fetch(input, {
        ...init,
        headers,
      });

      if (response.ok) {
        const contentType = response.headers.get("content-type") || "";
        if (contentType === "application/json") {
          return await response.json();
        }
        return response;
      } else {
        throw new Error("Unable to retrieve token");
      }
    } catch (error) {
      console.error("Fetch error:", error);
      throw error;
    }
  };
}
