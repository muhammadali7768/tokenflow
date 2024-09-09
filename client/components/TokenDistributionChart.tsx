"use client"
import 'chart.js/auto';
import { Pie } from 'react-chartjs-2';


const TokenDistributionChart = ({distribution}:any) => {
  
    const data = {
        labels: ['Admin Wallet', 'User Wallets', 'Staking Pool'],
        datasets: [
            {
                data: [parseFloat(distribution?.OwnerWallet),parseFloat(distribution?.UserWallet),parseFloat(distribution?.TotalStaked)],
                backgroundColor: ['#4CAF50', '#FFCE56', '#36A2EB'],
            },
        ],
    };

    return <div style={{ width: '100%', height: '500px' }}>
    <Pie data={data}  />
     </div>;
};

export default TokenDistributionChart;
