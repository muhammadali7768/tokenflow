"use client"
import { weiToEther } from '@/app/utils/helper';
import apiClient from '@/axios/apiClient';
import { useEffect, useState } from 'react';
import OverviewCard from './OverviewCard';

const StakedAmount = () => {
  const [stakedAmount, setStakedAmount] = useState<any>({});
  const [error, setError] = useState('');
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

  useEffect(() => {
    fetchStakedAmount();
  }, [user]);
  const fetchStakedAmount = async () => {
    if(user){
    try {
         const res = await apiClient.get(`/api/stake/${user?.wallet}`);        

      if (res.data.success) {
        setStakedAmount(res.data.stake);
        console.log("Stacke", res.data.stake)
      } else {
        setError('Failed to fetch staked amount');
      }
    } catch (err) {
      setError('Error fetching staked amount');
    }
}
  };

  return (
    // <div className="mt-4">
    //   {error ? (
    //     <p className="text-red-500">{error}</p>
    //   ) : (
    //     <p className="text-gray-700">Staked Amount: {weiToEther(stakedAmount?.Amount)} ENGC</p>
    //   )}
    // </div>
    <OverviewCard title="Staked Amount" value={`${weiToEther(stakedAmount?.Amount)} ENGC`} />
  );
};

export default StakedAmount;
