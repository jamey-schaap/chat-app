import { useEffect, useState } from "react";

const WebSocketEvent = {
	ChatMessageCreated: 0,
} as const;

type WebSocketEvent = (typeof WebSocketEvent)[keyof typeof WebSocketEvent];

type WebSocketEventData = {
	type: typeof WebSocketEvent.ChatMessageCreated;
	payload: { id: string };
};

let webSocket: WebSocket;
let reconnect = false;
export function useWebSocket() {
	useEffect(() => {
		if (!webSocket || reconnect) {
			webSocket = new WebSocket("ws://localhost:8080/ws");
			setupWebSocketMethods(webSocket);
			reconnect = false;
		}
	}, [reconnect]); // not sure if this works
}

function setupWebSocketMethods(socket: WebSocket) {
	socket.onopen = () => {
		console.log("Connected to the WebSocket server");
	};
	socket.onclose = () => {
		console.log("Disconnected from the WebSocket server");
		reconnect = true;
	};
	socket.onmessage = (event: MessageEvent<string>) => {
		console.log("Received from WebSocket server: " + event.data);
		const { type, payload } = JSON.parse(event.data) as WebSocketEventData;
		switch (type) {
			case WebSocketEvent.ChatMessageCreated:
				console.log(payload.id);
				break;
			default:
				console.log("wtf");
				break;
		}
	};
}
