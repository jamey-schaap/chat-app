import logoLight from "~/components/chat/logo-light.svg";
import logoDark from "~/components/chat/logo-dark.svg";
import SendIconSvg from "~/icons/send.icon.svg";

export const Chat = () => {
	return (
		<main className="flex items-center justify-center pt-16 pb-4">
			<div className="flex-1 flex flex-col items-center gap-16 min-h-0">
				<header className="flex flex-col items-center gap-9">
					<div className="w-[500px] max-w-[100vw] p-4">
						<img src={logoLight} alt="React Router" className="block w-full dark:hidden" />
						<img src={logoDark} alt="React Router" className="hidden w-full dark:block" />
					</div>
				</header>
				<div className="max-w-[300px] w-full space-y-6 px-4">
					<div className="rounded-3xl border border-gray-200 p-6 dark:border-gray-700 space-y-4">
						<p className="leading-6 text-gray-700 dark:text-gray-200 text-center">What&apos;s next?</p>
					</div>
					<div className="rounded-3xl border border-gray-200 p-6 dark:border-gray-700 space-y-4">
						<div className="flex items-center justify-center gap-4">
							<input
								type="text"
								placeholder="Type your message..."
								className="w-full rounded-3xl border border-gray-200 p-4 dark:border-gray-700"
							/>
							<button type="button" className="cursor-pointer">
								<SendIconSvg />
							</button>
						</div>
					</div>
				</div>
			</div>
		</main>
	);
};
