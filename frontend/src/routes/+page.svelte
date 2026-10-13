<script lang="ts">
	import { goto } from '$app/navigation';

	let username = $state('');
	let password = $state('');
	let role = $state('admin');

	let errorMessage = $state('');
	let isLoading = $state(false);

	async function handleLogin(event: SubmitEvent) {
		event.preventDefault();
		errorMessage = '';
		isLoading = true;

		try {
			const res = await fetch('http://localhost:8080/api/auth/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					username,
					password
				})
			});

			const data = await res.json();

			if (!res.ok) {
				throw new Error(data.error || 'Login gagal, periksa username dan password.');
			}

			// 1. Simpan Token JWT & Data User ke Browser
			localStorage.setItem('token', data.token);
			localStorage.setItem('user', JSON.stringify(data.user));

			// 2. Tembus Langsung ke Dashboard
			goto('/dashboard');
		} catch (err: any) {
			errorMessage = err.message;
		} finally {
			isLoading = false;
		}
	}
</script>

<div
	class="min-h-screen flex items-center justify-center bg-gradient-to-br from-gray-50 to-gray-200 px-4 sm:px-6 lg:px-8 font-sans selection:bg-gray-900 selection:text-white"
>
	<div class="max-w-md w-full">
		<!-- Area Logo & Sambutan -->
		<div class="text-center mb-10">
			<div
				class="inline-flex items-center justify-center w-16 h-16 rounded-2xl bg-gray-900 text-white mb-6 shadow-xl shadow-gray-900/20"
			>
				<svg
					class="w-8 h-8"
					fill="none"
					stroke="currentColor"
					viewBox="0 0 24 24"
					xmlns="http://www.w3.org/2000/svg"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M13 10V3L4 14h7v7l9-11h-7z"
					></path>
				</svg>
			</div>
			<h2 class="text-3xl font-bold text-gray-900 tracking-tight">Selamat Datang</h2>
			<p class="mt-2 text-sm text-gray-500 font-medium">
				Masuk ke dasbor sistem untuk melanjutkan
			</p>
		</div>

		<!-- Kartu Form Login -->
		<div
			class="bg-white/80 backdrop-blur-xl py-8 px-10 shadow-[0_8px_30px_rgb(0,0,0,0.06)] rounded-[2rem] border border-gray-100/80"
		>
			{#if errorMessage}
				<div
					class="mb-6 p-3.5 rounded-xl bg-red-50 border border-red-200 text-red-600 text-xs font-semibold text-center"
				>
					{errorMessage}
				</div>
			{/if}

			<form class="space-y-6" onsubmit={handleLogin}>
				<div class="space-y-5">
					<!-- Pilihan Role -->
					<div>
						<label class="block text-sm font-semibold text-gray-700 mb-2">Tipe Akses</label>
						<div class="relative">
							<select
								bind:value={role}
								class="appearance-none w-full bg-gray-50 border border-gray-200 text-gray-700 rounded-xl px-4 py-3.5 focus:outline-none focus:ring-2 focus:ring-gray-900 focus:border-transparent focus:bg-white transition-all duration-200 cursor-pointer"
							>
								<option value="admin">Perusahaan (Admin)</option>
								<option value="owner">PT Otak Kanan (Owner)</option>
							</select>
							<div
								class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-4 text-gray-400"
							>
								<svg
									class="h-4 w-4"
									fill="none"
									stroke="currentColor"
									viewBox="0 0 24 24"
									><path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M19 9l-7 7-7-7"
									></path></svg
								>
							</div>
						</div>
					</div>

					<!-- Input Username -->
					<div>
						<label class="block text-sm font-semibold text-gray-700 mb-2">Nama Pengguna</label>
						<input
							type="text"
							bind:value={username}
							required
							class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3.5 focus:outline-none focus:ring-2 focus:ring-gray-900 focus:border-transparent focus:bg-white transition-all duration-200 placeholder-gray-400"
							placeholder="Masukkan username"
						/>
					</div>

					<!-- Input Password -->
					<div>
						<label class="block text-sm font-semibold text-gray-700 mb-2">Kata Sandi</label>
						<input
							type="password"
							bind:value={password}
							required
							class="w-full bg-gray-50 border border-gray-200 rounded-xl px-4 py-3.5 focus:outline-none focus:ring-2 focus:ring-gray-900 focus:border-transparent focus:bg-white transition-all duration-200 placeholder-gray-400"
							placeholder="••••••••"
						/>
					</div>
				</div>

				<!-- Opsi Ekstra -->
				<div class="flex items-center justify-between mt-2">
					<div class="flex items-center">
						<input
							id="remember-me"
							type="checkbox"
							class="h-4 w-4 text-gray-900 focus:ring-gray-900 border-gray-300 rounded cursor-pointer accent-gray-900"
						/>
						<label for="remember-me" class="ml-2 block text-sm text-gray-600 cursor-pointer font-medium"
							>Ingat saya</label
						>
					</div>
					<div class="text-sm">
						<a
							href="/register"
							class="font-semibold text-gray-900 hover:text-gray-700 transition-colors"
							>Daftar Akun</a
						>
					</div>
				</div>

				<!-- Tombol Submit -->
				<button
					type="submit"
					disabled={isLoading}
					class="w-full flex justify-center py-3.5 px-4 mt-8 border border-transparent rounded-xl shadow-lg shadow-gray-900/20 text-sm font-bold text-white bg-gray-900 hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-900 transition-all duration-200 hover:-translate-y-0.5 active:translate-y-0 disabled:opacity-50"
				>
					{isLoading ? 'Memproses...' : 'Masuk Sekarang'}
				</button>
			</form>
		</div>
	</div>
</div>