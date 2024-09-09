import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { jwtVerify } from 'jose';


const SECRET_KEY  = new TextEncoder().encode('secret') //TODO: put the key in .env 

export async function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl;

  const token = request.cookies.get("token")?.value;

  if (!token) {
    return NextResponse.redirect(new URL("/auth/login", request.url));
  }

  try {
    const { payload } = await jwtVerify(token, SECRET_KEY);

    // Extract the role from the decoded token payload
    const userRole = payload.role;
   
      if (userRole === "owner") {
      if (pathname.startsWith("/admin")) {
        return NextResponse.next();
      } else {
        return NextResponse.redirect(new URL("/admin/dashboard", request.url));
      }
    } 
    else if(userRole=="user") {
      console.log("PATH",pathname)
      if(pathname=="/admin/deploy"){
     return NextResponse.redirect(new URL("/admin/dashboard", request.url));
    }
    else {
     return NextResponse.next();
    }
   }
    else {
      return NextResponse.redirect(new URL("/auth/login", request.url));
    }
  } catch (err) {
    console.error("JWT verification failed:", err);

    return NextResponse.redirect(new URL("/auth/login", request.url));
  }
}

export const config = {
  matcher: ["/admin/:path*", "/user/:path*"], // Apply to specific paths
};
