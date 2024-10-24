import {Link} from "react-router-dom";
import {SignedIn, SignedOut, UserButton} from "@clerk/clerk-react";

export default function Header() {
  return (
    <header className="p-4">
      <div className="flex items-center justify-between">
        <div>
          <Link to="/" className="font-semibold">Clerk demo</Link>
        </div>
        <SignedIn>
          <UserButton />
        </SignedIn>
        <SignedOut>
          <Link to="/sign-in">Sign in</Link>
        </SignedOut>
      </div>
      <nav>
        <SignedIn>
          <Link to="/profile">Profile</Link>
        </SignedIn>
      </nav>
    </header>
  );
}