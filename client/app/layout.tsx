import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css"

 const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Social Network App on Blockchain",
  description: "First Social network app on blockchain. Here you can earn crypto on the engagement",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
       <body className={inter.className}>{children}</body>  
     
    </html>

  );
}
