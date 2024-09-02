import * as React from 'react';
import { BarChart } from '@mui/x-charts/BarChart';

const uData = [4000, 3000, 2000, 2780, 1890, 2390, 3490];
const pData = [2400, 1398, 9800, 3908, 4800, 3800, 4300];
const xLabels = [
    'Page A',
    'Page B',
    'Page C',
    'Page D',
    'Page E',
    'Page F',
    'Page G',
];

export default function SimpleBarChart() {
    return (
        <div className={"grid grid-cols-3 justify-between "}>
            <div className={"flex flex-col items-start justify-center"}>
                <p>CPU</p>
                <BarChart
                    width={500}
                    height={300}
                    series={[
                        {data: pData, label: 'pv', id: 'pvId'},
                        {data: uData, label: 'uv', id: 'uvId'},
                    ]}
                    xAxis={[{data: xLabels, scaleType: 'band'}]}
                />

            </div>
            <div className={"flex flex-col items-start justify-center"}>
                <p>CPU</p>
                <BarChart
                    width={500}
                    height={300}

                    series={[
                        {data: pData, label: 'pv', id: 'pvId'},
                        {data: uData, label: 'uv', id: 'uvId'},
                    ]}
                    xAxis={[{data: xLabels, scaleType: 'band'}]}
                />

            </div>

        </div>
    );
}
