import axios from 'axios';

// Create an instance of Axios
const apiClient = axios.create({
    baseURL: `${process.env.NEXT_PUBLIC_BASE_API}/`, 
    withCredentials: true, 
    timeout: 10000,
});

// Add a request interceptor
apiClient.interceptors.request.use(
    (config) => {
        // Get the token from local storage
        const token = localStorage.getItem("token");
      
        // If the token exists, add it to the Authorization header
        // if (token) {
        //     config.headers['Authorization'] = `Bearer ${token}`;
        // }

        return config;
    },
    (error) => {
        // Handle the error
        return Promise.reject(error);
    }
);

// Optionally, add a response interceptor as well
apiClient.interceptors.response.use(
    (response) => {
        // Any status code that lies within the range of 2xx causes this function to trigger
        return response;
    },
    (error) => {
        // Any status codes that falls outside the range of 2xx causes this function to trigger
        return Promise.reject(error);
    }
);

export default apiClient;
