import React, { useEffect, useState } from "react";
import apiClient from "@/axios/apiClient";


const TransferReward = () => {
    const [recipientAddress, setRecipientAddress] = useState<string>("");
    const [amount, setAmount] = useState<string>();
    const [loading, setLoading] = useState<boolean>(false);
    const [engcBalance, setEngcBalance] = useState<number>(0);
    const handleTransfer = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        if (Number(amount) <= 0) {
            alert("Amount should be greater than zero");
            return;
        }

        setLoading(true);

        try {
            // Send request to the backend to process the token transfer
            const response = await apiClient.post(`api/send-reward`, {
                recipientAddr: recipientAddress,
                amount:Number(amount),
            });

            if (response.data.success) {
                alert("Reward transferred successfully!");
            } else {
                alert("Failed to transfer reward: " + response.data.message);
            }
        } catch (error) {
            alert("Error occurred while processing the transaction");
        } finally {
            setLoading(false);
        }
    };
    useEffect(() => {
        fetchEngcBalance()
    }, [])

    const fetchEngcBalance = async () => {
        const user = localStorage.getItem("user");
        if (user) {
            try {
                const parsedUser = JSON.parse(user);


                try {
                    const response = await apiClient.get(`api/engc-balance/${parsedUser.wallet}`);
                    let parsedVal = Number(Number(response.data.balance).toFixed(5));
                    setEngcBalance(parsedVal);
                } catch (err) {
                    console.error("Error fetching user data", err);

                }
            } catch (error) {
                console.error("Error parsing user from local storage:", error);
            }
        }
    };


    return (
        // <div className="max-w-lg mx-auto mt-10">
        <div className="w-full max-w-md p-8 bg-white rounded-lg shadow-md">
            <h2 className="text-2xl font-bold mb-6">Transfer ENGC Token</h2>
            <div className="mb-6">
                <p className="text-lg font-semibold">Balance:</p>
                <p className="text-gray-700">{engcBalance} ENGC</p>
            </div>
            <form onSubmit={handleTransfer} className="space-y-6">
                <div>
                    <label className="block text-gray-700">Recipient Address</label>
                    <input
                        type="text"
                        value={recipientAddress}
                        onChange={(e) => setRecipientAddress(e.target.value)}
                        placeholder="0x123..."
                        className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                </div>

                <div>
                    <label className="block text-gray-700">Amount (ENGC)</label>
                    <input
                        type="number"
                        value={amount}
                        onChange={(e) => setAmount(e.target.value)}
                        placeholder="Enter amount"
                        className="mt-1 block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                </div>

                <div>
                    <button
                        type="submit"
                        disabled={loading}
                        className={`w-full bg-blue-700 text-white font-bold py-2 px-4 rounded-lg  hover:bg-indigo-700 ${loading ? "opacity-50 cursor-not-allowed" : ""
                            }`}
                    >
                        {loading ? "Processing..." : "Transfer"}
                    </button>
                </div>
            </form>
        </div>
    );
};

export default TransferReward;
