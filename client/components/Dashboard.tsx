"use client"
import OverviewCard from '@/components/OverviewCard';
import TransactionTable from '@/components/TransactionTable';
import TokenDistributionChart from '@/components/TokenDistributionChart';
import { useEffect, useState } from 'react';
import apiClient from '@/axios/apiClient';
import StakedAmount from './StakedAmount';


const Dashboard = () => {
  const [address, setAddress] = useState<string | null>("");
  const [ethBalance, setEthBalance] = useState<number>(0);
  const [engcBalance, setEngcBalance] = useState<number>(0);
  const [transactions, setTransactions] = useState<any>();
  const [distribution,setDistribution]=useState<any>()
  const [amount, setAmount] = useState<number>(0);
  const [error, setError] = useState<string>("");
  useEffect(() => {
    fetchEthBalance();
    fetchEngcBalance();
    fetchTransactions()
  }, []);
  const fetchEthBalance = async () => {
    const user = localStorage.getItem("user");
    console.log("PARSED USER", user)
    if (user) {
      try {
        const parsedUser = JSON.parse(user);

        setAddress(parsedUser.wallet)
        try {
          const response = await apiClient.get(`api/balance/${parsedUser.wallet}`);
          let parsedVal = Number(Number(response.data.balance).toFixed(5));
          setEthBalance(parsedVal);
        } catch (err) {
          console.error("Error fetching user data", err);
          setError("Failed to load user data.");
        }
      } catch (error) {
        console.error("Error parsing user from local storage:", error);
      }
    }
  };
  const fetchTransactions = async () => {
 
    try {

      try {
        const response = await apiClient.get(`api/get-recent-transactions`);
        let tx:any= response.data.transactions;
        tx={...tx, TransferLogs: tx.TransferLogs || [], ApprovalLogs: tx.ApprovalLogs || []}
        setTransactions(tx);
      } catch (err) {
        console.error("Error fetching user data", err);
        setError("Failed to load user data.");
      }
    } catch (error) {
      console.error("Error parsing user from local storage:", error);
    }

  };

  const fetchEngcBalance = async () => {
    const user = localStorage.getItem("user");
    console.log("PARSED USER", user)
    if (user) {
      try {
        const parsedUser = JSON.parse(user);

        setAddress(parsedUser.wallet)
        try {
          const response = await apiClient.get(`api/engc-balance/${parsedUser.wallet}`);
          let parsedVal = Number(Number(response.data.balance).toFixed(5));
          setEngcBalance(parsedVal);
        } catch (err) {
          console.error("Error fetching user data", err);
          setError("Failed to load user data.");
        }
      } catch (error) {
        console.error("Error parsing user from local storage:", error);
      }
    }
  };


  useEffect(()=>{
   getDistribution()
  },[])
   const getDistribution=async()=>{
       let response= await apiClient.get("api/get-engc-token-distribution");
       
       setDistribution(response.data.distribution);
      }
  return (
    <div className="p-6 bg-gray-100 min-h-screen">
      <h1 className="text-3xl font-semibold text-gray-800 mb-4">Admin Dashboard <span className="text-sm">{address}</span></h1>

      {/* Overview Cards */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <OverviewCard title="ETH Balance" value={`${ethBalance} ETH`} />
        <OverviewCard title="Custom Coin Balance" value={`${engcBalance} ENGC`} />
        <OverviewCard title="Recent Transactions" value={`${transactions?.TransferLogs?.length + transactions?.ApprovalLogs?.length}`} />
        {/* <OverviewCard title="Staked Amount" value={`${engcBalance} ENGC`} /> */}
        <StakedAmount/>
      </div>

      {/* Token Distribution Chart */}
      <div className="bg-white p-4 rounded-lg shadow-md mb-8">
        <h2 className="text-xl font-semibold mb-4">Token Distribution <span className='text-sm text-red-300'>(Total Supply {parseFloat(distribution?.TotalSupply)})</span></h2>
        <TokenDistributionChart distribution={distribution} />
      </div>

      {/* Recent Transactions */}
      <div className="bg-white p-4 rounded-lg shadow-md">
        <h2 className="text-xl font-semibold mb-4">Recent Transactions</h2>
       { transactions?.TransferLogs?.length  > 0 &&
        <TransactionTable transactions={transactions} />
}
      </div>
    </div>
  );
};

export default Dashboard;
