interface OverviewCardProps {
    title: string;
    value: string;
  }
  
  const OverviewCard: React.FC<OverviewCardProps> = ({ title, value }) => {
    return (
      <div className="bg-white p-6 rounded-lg shadow-md">
        <h3 className="text-lg font-semibold text-gray-700 mb-2">{title}</h3>
        <p className="text-2xl font-bold text-gray-900">{value}</p>
      </div>
    );
  };
  
  export default OverviewCard;
  