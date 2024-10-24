import useFetch from "../hooks/useFetch.ts";
import {useEffect, useState} from "react";
import {SignedIn, SignedOut} from "@clerk/clerk-react";

export default function HomePage() {
  const fetcher = useFetch();
  const [privateDetails, setPrivateDetails] = useState();

  useEffect(() => {
    fetcher("http://localhost:42069/public").then((resp => console.log(resp)))
  }, []);

  async function getData() {
    try {
      const resp = await fetcher("http://localhost:42069/private");
      setPrivateDetails(resp);
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <>
      <p className="text-2xl mb-6">Home Page</p>
      <SignedIn>
        <div className="flex justify-between items-center">
          <p>You are signed in!</p>
          <button
            onClick={getData}
            className="bg-slate-900 text-white px-4 py-2 rounded shadow-lg hover:bg-slate-700 hover:shadow-none">
            Get details
          </button>
        </div>
        <pre className="text-sm bg-slate-200 p-4 rounded-lg shadow-lg my-4">
          {privateDetails
            ? JSON.stringify(privateDetails, null, 2)
            : "Click the button to fetch your details"}
        </pre>
      </SignedIn>
      <SignedOut>
        <p>Sign in to see more things...</p>
      </SignedOut>
    </>
  )
}
