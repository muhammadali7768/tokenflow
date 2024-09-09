"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import apiClient from "@/axios/apiClient";
import TransferReward from "@/components/TransferReward";


const TransferPage = () => {
    
    const [balance, setBalance] = useState<number>(0);
    const [recipientAddress, setRecipientAddress] = useState<string>("");
    const [amount, setAmount] = useState<string>();
    const [error, setError] = useState<string>("");
    const [success, setSuccess] = useState<string>("");

    const router = useRouter();


    useEffect(() => {
        fetchBalance();
    }, []);
    const fetchBalance = async () => {

        const user = localStorage.getItem("user");
        console.log("PARSED USER", user)
        if (user) {
            try {
                const parsedUser = JSON.parse(user);
                try {
                    const response = await apiClient.get(`api/balance/${parsedUser.wallet}`);
                    let parsedVal = Number(Number(response.data.balance).toFixed(5));
                    setBalance(parsedVal);
                } catch (err) {
                    console.error("Error fetching user data", err);
                    setError("Failed to load user data.");
                }
            } catch (error) {
                console.error("Error parsing user from local storage:", error);
            }
        }
    };

    const handleTransfer = async (e: React.FormEvent) => {
        e.preventDefault();

        try {
            const response = await apiClient.post("/api/send-balance", {
                recipientAddr: recipientAddress,
                amount: Number(amount)
            });

            if (response.data.success) {
                setSuccess("Transfer successful!");
                setError("");

                setTimeout(() => {
                    fetchBalance();
                }, 30000); // 30 seconds

            } else {
                setError(response.data.message);
                setSuccess("");
            }
        } catch (err: any) {
            console.log("ERROR", err)
            setError("Error: " + err.response.data.message);
            setSuccess("");
        }
    };

    return (
        // <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
        <div className="container mx-auto p-4">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div className="w-full max-w-md p-8 bg-white rounded-lg shadow-md">
                    <h1 className="text-2xl font-bold mb-4">Transfer ETH</h1>

                    <div className="mb-6">
                        <p className="text-lg font-semibold">Balance:</p>
                        <p className="text-gray-700">{balance} ETH</p>
                    </div>
                    <form onSubmit={handleTransfer} className="space-y-6">
                        <div className="mb-4">
                            <label className="block text-gray-700">Recipient Address</label>
                            <input
                                type="text"
                                id="recipientAddress"
                                value={recipientAddress}
                                placeholder="0x123..."
                                onChange={(e) => setRecipientAddress(e.target.value)}
                                className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                                required
                            />
                        </div>
                        <div className="mb-4">
                            <label htmlFor="amount" className="block text-gray-700">
                                Amount (ETH)
                            </label>
                            <input
                                type="number"
                                id="amount"
                                value={amount}
                                onChange={(e) => setAmount(e.target.value)}
                                className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                                required
                                placeholder="Enter amount"
                            />
                        </div>
                        <button
                            type="submit"
                            className="w-full py-2 px-4 bg-blue-700 text-white font-semibold rounded-md shadow-sm hover:bg-indigo-700"
                        >
                            Transfer
                        </button>
                        {error && <p className="mt-4 text-red-600">{error}</p>}
                        {success && <p className="mt-4 text-green-600">{success}</p>}
                    </form>
                </div>

                <TransferReward />
            </div>
        </div>
    );
};

export default TransferPage;
