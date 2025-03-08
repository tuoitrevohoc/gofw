interface AccessToken {
  accessToken: string;
  expiry: number;
}

export function handleLogin(viewer: AccessToken) {
  sessionStorage.setItem("accessToken", viewer.accessToken);
  sessionStorage.setItem("expiresAt", viewer.expiry.toString());

  console.log(viewer.accessToken, viewer.expiry);
}

export function handleLogout() {
  sessionStorage.removeItem("accessToken");
  sessionStorage.removeItem("expiresAt");

  location.href = "/auth/login";
}

export async function refreshToken() {
  const response = await fetch("/api/refresh-token", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });

  if (!response.ok) {
    return null;
  }

  const data: AccessToken = await response.json();
  handleLogin(data);
  return data.accessToken;
}

export async function getAccessToken() {
  const accessToken = sessionStorage.getItem("accessToken");
  const expiresAt = sessionStorage.getItem("expiresAt");

  console.log(accessToken, expiresAt);

  if (!accessToken || !expiresAt) {
    return null;
  }

  if (Date.now() > parseInt(expiresAt) * 1000) {
    return refreshToken();
  }

  return accessToken;
}
