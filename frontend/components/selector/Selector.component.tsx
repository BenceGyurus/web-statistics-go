import SelectorParams from "./SelectorParams";
import { Select } from 'antd';

export default function Selector({ options, selected, onSelect }:SelectorParams){
    return(
        <div>
            <Select style={{width : 300}} options={options.map(option=> {return {value : option, label: option}})} />
        </div>
    )
}