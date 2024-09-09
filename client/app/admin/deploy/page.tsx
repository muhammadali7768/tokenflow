"use client";
import { useState, useEffect } from 'react';
import axios from 'axios';
import { useRouter } from 'next/navigation';
import apiClient from '@/axios/apiClient';

interface User {
  username: string;
  role: string;
  wallet: string;
}

const AdminPage = () => {
  const [user, setUser] = useState<User | null>(null);
  const [contractAddress, setContractAddress] = useState<string | null>(null);
  const [error, setError] = useState<string>("");
  const [success, setSuccess] = useState<string>("");
  const router = useRouter();

  useEffect(() => {
    // Get the user details from local storage
    const storedUser = JSON.parse(localStorage.getItem('user') || '{}');
    if (storedUser && storedUser.wallet) {
      setUser(storedUser);
    } else {
      router.push('/auth/login'); // Redirect to login if no user found
    }
  }, [router]);

  const deployContract = async () => {
    try {
      const response = await apiClient.post('/api/deploy-engage-coin-contract');
      setContractAddress(response.data.contract_address);
       setSuccess(response.data.message)
    } catch (error:any) {
      console.error('Error deploying contract: ', error);
     setError('Error deploying contract: ' + error.response.data.message)
    }
  };

  const deployEngStackingContract = async () => {
    try {
      const response = await apiClient.post('/api/deploy-engc-stacking-contract');
      setContractAddress(response.data.contract_address);
       setSuccess(response.data.message)
    } catch (error:any) {
      console.error('Error deploying contract: ', error);
     setError('Error deploying contract: ' + error.response.data.message)
    }
  };

  if (!user) {
    return <div>Loading...</div>;
  }

  return (
    <div className="min-h-screen bg-gray-100 flex flex-col items-center justify-center">
      <div className="bg-white p-6 rounded-lg shadow-lg w-full max-w-md">
        <h1 className="text-2xl font-semibold text-center mb-4">User's Detail</h1>
        <div className="mb-4">
          <label className="block font-medium text-gray-700">User Name:</label>
          <p className="text-gray-800">{user.username}</p>
        </div>
        <div className="mb-4">
          <label className="block font-medium text-gray-700">Role:</label>
          <p className="text-gray-800">{user.role}</p>
        </div>
        <div className="mb-4">
          <label className="block font-medium text-gray-700">Wallet Address:</label>
          <p className="text-gray-800">{user.wallet}</p>
        </div>
        {contractAddress && (
          <div className="mb-4">
            <label className="block font-medium text-gray-700">Contract Address:</label>
            <p className="text-gray-800">{contractAddress}</p>
          </div>
        )}
        <button
          onClick={deployContract}
          className="w-full bg-blue-600 text-white font-semibold py-2 px-4 rounded-lg hover:bg-blue-700"
        >
          Deploy ENGC Contract
        </button>

        <button
          onClick={deployEngStackingContract}
          className="w-full bg-blue-600 text-white font-semibold py-2 px-4 rounded-lg hover:bg-blue-700 mt-5"
        >
          Deploy ENGC Stacking Contract
        </button>

        {error && <p className="mt-4 text-red-600">{error}</p>}
        {success && <p className="mt-4 text-green-600">{success}</p>}
      </div>
    </div>
  );
};

export default AdminPage;
