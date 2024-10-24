import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import RootLayout from "./layouts/root-layout.tsx";
import HomePage from "./pages/HomePage.tsx";
import SignInPage from "./pages/clerk/SignInPage.tsx";
import ProfilePage from "./pages/clerk/ProfilePage.tsx";

const router = createBrowserRouter([
  {
    element: <RootLayout />,
    children: [
      { path: "/", element: <HomePage /> },
      { path: "/sign-in/*", element: <SignInPage /> },
      { path: "/profile", element: <ProfilePage /> },
    ]
  }
])

createRoot(document.getElementById("root")!).render(
  <StrictMode>
      <RouterProvider router={router} />
  </StrictMode>,
)
