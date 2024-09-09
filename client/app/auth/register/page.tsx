"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import axios from "axios";
import apiClient from "@/axios/apiClient";
import Link from "next/link";

const Register = () => {
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [success, setSuccess] = useState("");

    const router = useRouter();

    const handleRegister = async (e: React.FormEvent) => {
        e.preventDefault();

        try {
            const response = await apiClient.post("/api/auth/register", {
                username,
                email,
                password
            });

            if (response.data.success) {
                setSuccess("Registration successful!");
                setError("");
                router.push("/auth/register");
            } else {
                setError(response.data.message);
                setSuccess("");
            }
        } catch (err) {
            setError("An error occurred during registration.");
            setSuccess("");
        }
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <div className="w-full max-w-md p-8 bg-white rounded-lg shadow-md">
                <h1 className="text-2xl font-bold mb-4">Register</h1>
                <form onSubmit={handleRegister}>
                    <div className="mb-4">
                        <label htmlFor="username" className="block text-sm font-medium text-gray-700">
                            Username
                        </label>
                        <input
                            type="text"
                            id="username"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                            required
                        />
                    </div>
                    <div className="mb-4">
                        <label htmlFor="email" className="block text-sm font-medium text-gray-700">
                            Email
                        </label>
                        <input
                            type="email"
                            id="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                            required
                        />
                    </div>
                    <div className="mb-6">
                        <label htmlFor="password" className="block text-sm font-medium text-gray-700">
                            Password
                        </label>
                        <input
                            type="password"
                            id="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                            required
                        />
                    </div>
                    <button
                        type="submit"
                        className="w-full py-2 px-4 bg-indigo-600 text-white font-semibold rounded-md shadow-sm hover:bg-indigo-700"
                    >
                        Register
                    </button>
                    {error && <p className="mt-4 text-red-600">{error}</p>}
                    {success && <p className="mt-4 text-green-600">{success}</p>}

                    <div className='mt-5'>Already have account <Link href="/auth/login" className='text-center text-lg text-blue-400'>Login here</Link></div>
                </form>
            </div>
        </div>
    );
};

export default Register;
