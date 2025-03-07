import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

export default function Home() {
  const navigate  = useNavigate();

  useEffect(() => {
    navigate("/auth/register");
  }, [navigate]);

  return (
    <div>
      <h1>Home</h1>
    </div>
  )
}