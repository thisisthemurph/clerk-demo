import {SignIn, useUser} from "@clerk/clerk-react";
import {useNavigate} from "react-router-dom";

export default function SignInPage() {
  const navigate = useNavigate();
  const {user} = useUser();

  if (!user) {
    return <SignIn path="/sign-in" forceRedirectUrl="/" />
  }

  navigate("/");
}