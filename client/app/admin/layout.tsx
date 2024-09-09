
import Header from "@/components/layout/Header";
import Sidebar from "@/components/layout/Sidebar";

export default function AdminLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    
      <div className="flex h-screen">
    <Sidebar />
    <div className="flex-1 flex flex-col">
      <Header />
      <main className="p-2 overflow-auto flex-1">{children}</main>
    </div>
  </div>
     

  );
}
