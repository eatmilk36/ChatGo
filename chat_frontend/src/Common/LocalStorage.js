export const setToken = (token) => {
    localStorage.setItem('token', token);
};

// Function to get token from localStorage
export const getToken = () => {
    return localStorage.getItem('token');
};

// Function to remove token from localStorage
export const removeToken = () => {
    localStorage.removeItem('token');
};