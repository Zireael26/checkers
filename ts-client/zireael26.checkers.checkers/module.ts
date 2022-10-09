// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgRejectGame } from "./types/checkers/tx";
import { MsgCreateGame } from "./types/checkers/tx";
import { MsgPlayMove } from "./types/checkers/tx";


export { MsgRejectGame, MsgCreateGame, MsgPlayMove };

type sendMsgRejectGameParams = {
  value: MsgRejectGame,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateGameParams = {
  value: MsgCreateGame,
  fee?: StdFee,
  memo?: string
};

type sendMsgPlayMoveParams = {
  value: MsgPlayMove,
  fee?: StdFee,
  memo?: string
};


type msgRejectGameParams = {
  value: MsgRejectGame,
};

type msgCreateGameParams = {
  value: MsgCreateGame,
};

type msgPlayMoveParams = {
  value: MsgPlayMove,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgRejectGame({ value, fee, memo }: sendMsgRejectGameParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRejectGame: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRejectGame({ value: MsgRejectGame.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRejectGame: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateGame({ value, fee, memo }: sendMsgCreateGameParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateGame: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateGame({ value: MsgCreateGame.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateGame: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgPlayMove({ value, fee, memo }: sendMsgPlayMoveParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgPlayMove: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgPlayMove({ value: MsgPlayMove.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgPlayMove: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgRejectGame({ value }: msgRejectGameParams): EncodeObject {
			try {
				return { typeUrl: "/zireael26.checkers.checkers.MsgRejectGame", value: MsgRejectGame.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRejectGame: Could not create message: ' + e.message)
			}
		},
		
		msgCreateGame({ value }: msgCreateGameParams): EncodeObject {
			try {
				return { typeUrl: "/zireael26.checkers.checkers.MsgCreateGame", value: MsgCreateGame.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateGame: Could not create message: ' + e.message)
			}
		},
		
		msgPlayMove({ value }: msgPlayMoveParams): EncodeObject {
			try {
				return { typeUrl: "/zireael26.checkers.checkers.MsgPlayMove", value: MsgPlayMove.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgPlayMove: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]>;

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });
		this.tx = txClient({ signer: client.signer, addr: client.env.rpcURL, prefix: client.env.prefix ?? "cosmos" });
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			Zireael26CheckersCheckers: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;