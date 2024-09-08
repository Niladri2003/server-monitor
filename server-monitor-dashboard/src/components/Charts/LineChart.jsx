import React, { useState, useEffect } from 'react';
import Stack from '@mui/material/Stack';
import { LineChart } from '@mui/x-charts/LineChart';
import axios from "axios";
import {CircularProgress, FormControl, InputLabel, MenuItem, Select, Typography} from "@mui/material";


const timeRanges = [
    { label: 'Last 1 hour', start: '-1h', stop: 'now()' },
    { label: 'Last 6 hours', start: '-6h', stop: 'now()' },
    { label: 'Last 12 hours', start: '-12h', stop: 'now()' },
    { label: 'Last 24 hours', start: '-24h', stop: 'now()' },
    { label: 'Last 7 days', start: '-7d', stop: 'now()' },
    { label: 'Custom Range', start: '', stop: '' }, // This can be used for custom time picker
];


export default function LineChartConnectNulls() {
    const [data, setData] = useState([]);
    const [xData, setXData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [selectedRange, setSelectedRange] = useState(timeRanges[1]);

    // Function to fetch API data
    const fetchData = async (start, stop) => {
        setLoading(true);
        try {

            const response = await axios.get('http://127.0.0.1:5000/api/v1/server/disk-usage',{ params: { start, stop } });
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
        setLoading(false);
    };

    // Fetch data on component mount
    useEffect(() => {
        if (selectedRange.start && selectedRange.stop) {
            fetchData(selectedRange.start, selectedRange.stop);
        }
    }, [selectedRange]);
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
        <Stack sx={{ width: '100%',alignItems: 'flex-end'  }}>
            <FormControl size={"small"} sx={{width:'30%' ,mr:'20px'}}>
                {/*<InputLabel id="time-range-label">Time Range</InputLabel>*/}
                <Select
                    labelId="time-range-label"
                    value={selectedRange.label}
                    onChange={(e) => {
                        const selected = timeRanges.find((range) => range.label === e.target.value);
                        setSelectedRange(selected);
                    }}
                >
                    {timeRanges.map((range, index) => (
                        <MenuItem key={index} value={range.label}>
                            {range.label}
                        </MenuItem>
                    ))}
                </Select>
            </FormControl>
            <LineChart
                xAxis={[{ data: xData, scaleType: 'point' }]}
                series={[{ data, connectNulls: true }]}
                height={250}
                margin={{ top: 10, bottom: 20 }}
            />
        </Stack>
    );
}
