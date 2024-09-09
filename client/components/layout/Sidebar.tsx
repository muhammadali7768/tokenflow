"use client"
import React, { useEffect, useState } from 'react';
import Link from 'next/link';

const Sidebar = () => {
    const [user, setUser] = useState<any>(null)
    useEffect(() => {
        const userStored = localStorage.getItem("user");
        if (userStored) {
            try {
                setUser(JSON.parse(userStored));
            }
            catch (error) {
                console.error("Error parsing user from local storage:", error);
            }
        }
    }, [])
  return (
    <div className="h-screen bg-gray-800 text-white w-64 flex flex-col space-y-4 p-4">
      <h2 className="text-2xl font-bold mb-8">Menu</h2>
      <Link href="/admin/dashboard" className="p-2 hover:bg-gray-700 rounded">
        Dashboard
      </Link>
      <Link href="/admin/transfer" className="p-2 hover:bg-gray-700 rounded">
        Transfer
      </Link>
     
      { user?.role=="owner" &&
      <Link href="/admin/deploy" className="p-2 hover:bg-gray-700 rounded">
        Deploy
      </Link>
}
<Link href="/admin/stake" className="p-2 hover:bg-gray-700 rounded">
        Stack
      </Link>
      <Link href="/posts" className="p-2 hover:bg-gray-700 rounded">
        Posts
      </Link>
    </div>
  );
};

export default Sidebar;
