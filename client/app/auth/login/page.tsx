"use client";
import apiClient from '@/axios/apiClient';
import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';

export default function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const router = useRouter();

  const handleSubmit = async (e:any) => {
    e.preventDefault();
    try {
      const response = await apiClient.post(`api/auth/login`, { email, password });
      console.log("REsponse",response.data)
      if (response.data.success) {
       console.log("success",response.data.success)
        localStorage.setItem('token', response.data.token);
        localStorage.setItem('user', JSON.stringify(response.data.user));
        router.push('/admin/dashboard');
      }
    } catch (error) {
      console.error('Login failed', error);
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <form onSubmit={handleSubmit} className="bg-white p-8 rounded-lg shadow-lg">
        <h2 className="text-2xl font-bold mb-6">Login</h2>
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className="w-full p-2 mb-4 border border-gray-300 rounded"
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="w-full p-2 mb-4 border border-gray-300 rounded"
        />
        <button
          type="submit"
          className="w-full p-2 bg-blue-500 text-white font-bold rounded hover:bg-blue-600"
        >
          Login
        </button>
        <div className='mt-5'>Don't have account <Link href="/auth/register" className='text-center text-lg text-blue-400'>Register here</Link></div>
      </form>
    </div>
  );
}
