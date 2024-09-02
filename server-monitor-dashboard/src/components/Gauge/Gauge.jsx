
import { Gauge, gaugeClasses } from '@mui/x-charts/Gauge'

const Gaugedemo = () => {
    return (
        <div className={"grid grid-cols-3 justify-between "}>
            <div className={"flex flex-col items-start justify-center "}>
                <div className={"h-28"}>
                    <Gauge
                        value={75}
                        startAngle={-110}
                        endAngle={110}
                        sx={{
                            [`& .${gaugeClasses.valueText}`]: {
                                fontSize: 28,
                                transform: 'translate(0px, 0px)',
                            },
                        }}
                        text={
                            ({value, valueMax}) => `${value} / ${valueMax}`
                        }
                    />
                </div>
            </div>
        </div>

                )
                }
                export default Gaugedemo