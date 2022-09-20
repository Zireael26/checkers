import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateGame } from "./types/checkers/tx";
import { MsgPlayMove } from "./types/checkers/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zireael26.checkers.checkers.MsgCreateGame", MsgCreateGame],
    ["/zireael26.checkers.checkers.MsgPlayMove", MsgPlayMove],
    
];

export { msgTypes }