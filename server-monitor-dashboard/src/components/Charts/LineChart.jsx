import React, { useState, useEffect } from 'react';
import Stack from '@mui/material/Stack';
import { LineChart } from '@mui/x-charts/LineChart';
import axios from "axios";
import {CircularProgress, Typography} from "@mui/material";

const data = [4000, 3000, 2000, null, 1890, 2390, 3490];
const xData = ['Page A', 'Page B', 'Page C', 'Page D', 'Page E', 'Page F', 'Page G'];
export default function LineChartConnectNulls() {
    const [data, setData] = useState([]);
    const [xData, setXData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    // Function to fetch API data
    const fetchData = async () => {
        try {
            const response = await axios.get('http://127.0.0.1:5000/api/v1/server/disk-usage');
            const apiData = response.data;

            // Extract used_gb values and times for the chart
            const usedGbData = apiData.map(item => item.used_gb);
            const timeLabels = apiData.map(item =>
                new Date(item.time).toLocaleTimeString()
            );

            setData(usedGbData);
            setXData(timeLabels);
        } catch (err) {
            setError('Failed to fetch data');
        } finally {
            setLoading(false);
        }
    };

    // Fetch data on component mount
    useEffect(() => {
        fetchData();
    }, []);
    // Show loader while data is being fetched
    if (loading) {
        return (
            <Stack
                sx={{ width: '100%', height: '200px' }}
                justifyContent="center"
                alignItems="center"
            >
                <CircularProgress />
                <Typography>Loading data...</Typography>
            </Stack>
        );
    }
    // Show error message if API call fails
    if (error) {
        return (
            <Stack
                sx={{ width: '100%', height: '200px' }}
                justifyContent="center"
                alignItems="center"
            >
                <Typography color="error">{error}</Typography>
            </Stack>
        );
    }

    return (
        <Stack sx={{ width: '100%' }}>

            <LineChart
                xAxis={[{ data: xData, scaleType: 'point' }]}
                series={[{ data, connectNulls: true }]}
                height={200}
                margin={{ top: 10, bottom: 20 }}
            />
        </Stack>
    );
}
