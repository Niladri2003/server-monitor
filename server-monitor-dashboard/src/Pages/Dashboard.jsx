import SimpleBarChart from "../components/Charts/BarChart.jsx";
import BasicLineChart from "../components/Charts/LineChart.jsx";
import Gaugedemo from "../components/Gauge/Gauge.jsx";
import {Gauge, gaugeClasses} from "@mui/x-charts/Gauge";



const Dashboard = () => {
    return (
        <div className="flex flex-1 h-full">
            <div
                className="md:p-4 flex flex-col gap-2 flex-1 w-full h-full">
                <div className={"grid grid-cols-7 mb-5"}>
                    <div className={"h-36 flex items-center flex-col"}>
                        <p className={"font-semibold"}>CPU Busy</p>
                        <Gauge
                            innerRadius="73%"
                            value={75}
                            startAngle={-110}
                            endAngle={110}
                            sx={{
                                [`& .${gaugeClasses.valueText}`]: {
                                    fontSize: 30,
                                    transform: 'translate(0px, 0px)',
                                },
                            }}
                            text={
                                ({value}) => `${value}%`
                            }
                        />

                    </div>
                    <div className={"h-36 flex items-center flex-col"}>
                        <p className={"font-semibold"}>Sys Load</p>
                        <Gauge
                            innerRadius="73%"
                            value={75}
                            startAngle={-110}
                            endAngle={110}
                            sx={{
                                [`& .${gaugeClasses.valueText}`]: {
                                    fontSize: 30,
                                    transform: 'translate(0px, 0px)',
                                },
                            }}
                            text={
                                ({value}) => `${value}`
                            }
                        />

                    </div>
                    <div className={"h-36 flex items-center flex-col"}>
                        <p className={"font-semibold"}>Ram Used</p>
                        <Gauge
                            innerRadius="73%"
                            value={75}
                            startAngle={-110}
                            endAngle={110}
                            sx={{
                                [`& .${gaugeClasses.valueText}`]: {
                                    fontSize: 30,
                                    transform: 'translate(0px, 0px)',
                                },
                            }}
                            text={
                                ({value}) => `${value}`
                            }
                        />

                    </div>
                    <div className={"h-36 flex items-center flex-col"}>
                        <p className={"font-semibold"}>Swap used</p>
                        <Gauge
                            innerRadius="73%"
                            value={75}
                            startAngle={-110}
                            endAngle={110}
                            sx={{
                                [`& .${gaugeClasses.valueText}`]: {
                                    fontSize: 30,
                                    transform: 'translate(0px, 0px)',
                                },
                            }}
                            text={
                                ({value}) => `${value}`
                            }
                        />

                    </div>
                    <div className={"h-36 flex items-center flex-col"}>
                        <p className={"font-semibold"}>Swap used</p>
                        <Gauge
                            innerRadius="73%"
                            value={75}
                            startAngle={-110}
                            endAngle={110}
                            sx={{
                                [`& .${gaugeClasses.valueText}`]: {
                                    fontSize: 30,
                                    transform: 'translate(0px, 0px)',
                                },
                            }}
                            text={
                                ({value}) => `${value}`
                            }
                        />

                    </div>
                    <div className={"flex flex-col items-center justify-center gap-4"}>
                        <div className={"flex flex-col items-center justify-center"}>
                            <div>Cpu cores</div>
                            <div>2</div>
                        </div>
                        <div className={"flex flex-col items-center justify-center"}>
                            <div>Root Fs used</div>
                            <div>2</div>
                        </div>
                    </div>
                    <div className={"flex flex-col items-center justify-center gap-4 p-2"}>
                        <div className={"flex flex-col items-center justify-center"}>
                            <div>Uptime</div>
                            <div>2.8 days</div>
                        </div>
                        <div className={"flex flex-row items-center justify-between w-full"}>
                            <div className={"flex flex-col items-center justify-center"}>
                                <div>Ram Total</div>
                                <div>2</div>
                            </div>
                            <div className={"flex flex-col items-center justify-center"}>
                                <div>Swap Total</div>
                                <div>2</div>
                            </div>
                        </div>
                    </div>
                </div>


                <div className={"grid grid-cols-3 gap-2 space-y-3 pb-2 mt-10"}>
                    <div className={"w-[100%] flex flex-col items-center"}>
                        <p>Cpu Usage</p>
                        <BasicLineChart/>
                    </div>
                    <div className={"w-[100%] flex flex-col items-center"}>
                        <p>Memory Usage</p>
                        <BasicLineChart/>
                    </div>
                    <div className={"w-[100%] flex flex-col items-center"}>
                        <p>Disk Usage</p>
                        <BasicLineChart/>
                    </div>

                </div>
                <div>
                    <p className={"lg:text-xl"}>Memory Info</p>
                    <div className={"mt-10"}>
                        <SimpleBarChart/>
                    </div>
                </div>


            </div>

        </div>
    )
}
export default Dashboard;