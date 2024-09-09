"use client"
import apiClient from '@/axios/apiClient';
import { useState } from 'react';
import StakedAmount from '@/components/StakedAmount';
const StakePage = () => {
  const [amount, setAmount] = useState('');
  const [error, setError] = useState('');
  const [message, setMessage] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setMessage('');

    // API call to backend Go Fiber server
    try {
      const res = await apiClient.post('/api/stake-engc', { amount: parseFloat(amount) });

      if (res.data.success) {
        setMessage(`Successfully staked ${amount} ENGC tokens!`);
        setAmount('');
      } else {
        setError(res.data.message || 'Something went wrong');
      }
    } catch (err:any) {
      setError(err.response.data.message || 'Something went wrong');
    }
  };

  return (
    <div className="min-h-screen flex justify-center items-center bg-gray-100">
      <div className="bg-white p-6 rounded-lg shadow-lg w-full max-w-md">
        <h1 className="text-2xl font-bold mb-4">Stake ENGC Tokens</h1>
        {error && <p className="text-red-500">{error}</p>}
        {message && <p className="text-green-500">{message}</p>}

        {/* <StakedAmount /> */}
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label htmlFor="amount" className="block text-sm font-medium text-gray-700">
              Amount (ENGC)
            </label>
            <input
              type="number"
              id="amount"
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              value={amount}
              onChange={(e) => setAmount(e.target.value)}
              required
            />
          </div>
          <button
            type="submit"
            className="w-full bg-indigo-500 text-white py-2 px-4 rounded hover:bg-indigo-600 focus:outline-none"
          >
            Stake Tokens
          </button>
        </form>
      </div>
    </div>
  );
};

export default StakePage;
