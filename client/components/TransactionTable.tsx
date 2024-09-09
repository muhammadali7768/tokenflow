
interface MyComponentProps {
  transactions: any;
}
import { weiToEther } from "@/app/utils/helper";
const TransactionTable = ({transactions}: MyComponentProps) => {
    // const transactions = [
    //   { id: 1, type: 'ETH Transfer', amount: '2.0 ETH', status: 'Success' },
    //   { id: 2, type: 'Custom Coin Mint', amount: '10,000 ENGC', status: 'Pending' },
    //   { id: 3, type: 'ETH Transfer', amount: '1.5 ETH', status: 'Failed' },
    // ];

  
  
    return (
      <table className="min-w-full bg-white">
        <thead>
          <tr>
            <th className="text-left py-2 px-4 font-medium text-gray-700">Type</th>
            <th className="text-left py-2 px-4 font-medium text-gray-700">Amount</th>
            <th className="text-left py-2 px-4 font-medium text-gray-700">From</th>
            <th className="text-left py-2 px-4 font-medium text-gray-700">To</th>
            <th className="text-left py-2 px-4 font-medium text-gray-700">Status</th>
          </tr>
        </thead>
        <tbody>
          {transactions.TransferLogs.map((transaction:any, key:any) => (
            <tr key={key}>
              <td className="py-2 px-4">Transfer</td>
              <td className="py-2 px-4">{weiToEther(transaction.Value)}</td>
              <td className="py-2 px-4">{transaction.From}</td>
              <td className="py-2 px-4">{transaction.To}</td>
              <td className={`py-2 px-4 ${transaction.Status === 'completed' ? 'text-green-500' : 'text-red-500'}`}>
                {transaction.Status}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    );
  };
  
  export default TransactionTable;
  