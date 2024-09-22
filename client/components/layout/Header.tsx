"use client"
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import React, { useEffect, useState } from 'react';
import Cookies from 'js-cookie';
const Header = () => {
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

    const router = useRouter();

    const shortWalletAddress = `${user?.wallet?.slice(0, 4)}...${user?.wallet?.slice(-4)}`;

    const copyToClipboard = () => {
        navigator.clipboard.writeText(user.wallet);
        alert('Wallet address copied!');
    };

    const handleLogout = () => {
        // Clear cookies (JWT token)
        Cookies.remove('token');

        // Clear localStorage (if you store any data there)
        localStorage.clear();

        // Redirect to login page
        router.push('/auth/login');
    };

    return (
        <header className="bg-gray-900 text-white h-16 flex items-center justify-between px-8">
            <h1 className="text-2xl font-bold">TokenFlow Blockchain</h1>
            <div>
                <input
                    type="text"
                    placeholder="Search..."
                    className="px-4 py-2 rounded bg-gray-700 text-white focus:outline-none"
                />
            </div>
            <div className="flex space-x-4">

                {
                    !user ? (
                        <>
                            <Link href="/" className="bg-blue-500 px-4 py-2 rounded hover:bg-blue-600">Login</Link>
                            <Link href="/register" className="bg-green-500 px-4 py-2 rounded hover:bg-green-600">Register</Link>
                        </>
                    )
                        : <>
                            <div className="text-right">
                                <p className="text-sm text-white">{user.username}</p>
                                <span className="text-xs text-white flex items-center">{shortWalletAddress}

                                    <button onClick={copyToClipboard} className="ml-2 p-1 bg-gray-700 rounded hover:bg-gray-600">
                                        {/* <svg xmlns="http://www.w3.org/2000/svg" className="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M8 16h8M8 12h8m-8 4h8m4-8h-8m-8-4H5m4-2h8M5 12H3m16-8h2m-10 12v2h8v-2M9 8h4" />
              </svg> */}
                                        <svg fill="#FFFFFF" className="h-4 w-4" viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" stroke="#FFFFFF"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> 
                                            <path d="M47.81 91.725c0-8.328 6.539-15.315 15.568-15.33 9.03-.016 14.863.015 14.863.015s-.388-8.9-.388-15.978c0-7.08 6.227-14.165 15.262-14.165s92.802-.26 101.297.37c8.495.63 15.256 5.973 15.256 14.567 0 8.594-.054 93.807-.054 101.7 0 7.892-7.08 15.063-15.858 15.162-8.778.1-14.727-.1-14.727-.1s.323 9.97.323 16.094c0 6.123-7.12 15.016-15.474 15.016s-93.117.542-101.205.542c-8.088 0-15.552-7.116-15.207-15.987.345-8.871.345-93.58.345-101.906zm46.06-28.487l-.068 98.164c0 1.096.894 1.99 1.999 1.984l95.555-.51a2.007 2.007 0 0 0 1.998-2.01l-.064-97.283a2.01 2.01 0 0 0-2.01-2.007l-95.4-.326a1.99 1.99 0 0 0-2.01 1.988zM63.268 95.795l.916 96.246a2.007 2.007 0 0 0 2.02 1.982l94.125-.715a3.976 3.976 0 0 0 3.953-4.026l-.137-11.137s-62.877.578-71.054.578-15.438-7.74-15.438-16.45c0-8.71.588-68.7.588-68.7.01-1.1-.874-1.99-1.976-1.975l-9.027.13a4.025 4.025 0 0 0-3.97 4.067z" fill-rule="evenodd"></path> </g></svg>
                                    </button>
                                </span>
                            </div>
                            <button onClick={handleLogout} className="bg-red-500 px-4 py-2 rounded hover:bg-red-800">Logout</button>

                        </>
                }



            </div>
        </header>
    );
};

export default Header;
