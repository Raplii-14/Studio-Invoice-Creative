<script lang="ts">
	import { onMount } from 'svelte';

	let showModalLocal = $state(false);
	let isLoading = $state(true);

	// State Data dari Backend
	let stats = $state({
		total_pendapatan: 0,
		menunggu_pembayaran: { total: 0, count: 0 },
		jatuh_tempo: { total: 0, count: 0 }
	});
	let invoices = $state<any[]>([]);
	let clients = $state<any[]>([]);

	// State Form Modal
	let nomorInvoice = $state('');
	let selectedClientId = $state('');
	let tanggalTerbit = $state(new Date().toISOString().split('T')[0]);
	let tanggalJatuhTempo = $state('');
	let deskripsi = $state('');
	let jumlahTagihan = $state(0);
	let isSubmitting = $state(false);

	// Helper Format Rupiah
	function formatRupiah(val: number) {
		return new Intl.NumberFormat('id-ID', {
			style: 'currency',
			currency: 'IDR',
			maximumFractionDigits: 0
		}).format(val || 0);
	}

	// Helper Styling Status
	function getStatusStyle(status: string) {
		if (status === 'Lunas') return 'bg-emerald-50 text-emerald-600 border border-emerald-200/60';
		if (status === 'Belum Bayar') return 'bg-amber-50 text-amber-600 border border-amber-200/60';
		return 'bg-rose-50 text-rose-600 border border-rose-200/60';
	}

	// Load Data dari Backend Pas Halaman Dibuka
	async function loadDashboardData() {
		const token = localStorage.getItem('token');
		if (!token) {
			window.location.href = '/login';
			return;
		}

		const headers = { Authorization: `Bearer ${token}` };

		try {
			// Fetch Stats, Invoices, dan Clients secara paralel
			const [resStats, resInvoices, resClients] = await Promise.all([
				fetch('http://localhost:8080/api/dashboard/stats', { headers }),
				fetch('http://localhost:8080/api/invoices', { headers }),
				fetch('http://localhost:8080/api/clients', { headers })
			]);

			if (resStats.ok) stats = await resStats.json();
			if (resInvoices.ok) invoices = await resInvoices.json();
			if (resClients.ok) clients = await resClients.json();
		} catch (err) {
			console.error('Gagal mengambil data dari backend:', err);
		} finally {
			isLoading = false;
		}
	}

	onMount(() => {
		loadDashboardData();
	});

	// Handle Simpan Tagihan Baru
	async function handleSubmitLocal(event: SubmitEvent) {
		event.preventDefault();
		const token = localStorage.getItem('token');
		if (!token) return;

		isSubmitting = true;

		try {
			const res = await fetch('http://localhost:8080/api/invoices', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${token}`
				},
				body: JSON.stringify({
					nomor_invoice: nomorInvoice,
					client_id: Number(selectedClientId),
					tanggal_terbit: tanggalTerbit,
					tanggal_jatuh_tempo: tanggalJatuhTempo,
					deskripsi: deskripsi,
					items: [
						{
							nama_item: deskripsi || 'Tagihan Layanan',
							quantity: 1,
							unit_price: Number(jumlahTagihan)
						}
					]
				})
			});

			if (!res.ok) {
				const errData = await res.json();
				alert('Gagal menyimpan tagihan: ' + (errData.error || 'Error server'));
				return;
			}

			// Reset Form & Reload Data
			showModalLocal = false;
			deskripsi = '';
			jumlahTagihan = 0;
			selectedClientId = '';
			loadDashboardData();
		} catch (err) {
			alert('Terjadi kesalahan jaringan');
		} finally {
			isSubmitting = false;
		}
	}
</script>

<!-- Header Halaman Dashboard -->
<header
	class="h-24 flex items-center justify-between px-10 bg-white/70 backdrop-blur-md sticky top-0 z-10 border-b border-slate-100/80"
>
	<div>
		<h1 class="text-2xl font-black text-slate-900 tracking-tight">Dashboard</h1>
		<p class="text-sm text-slate-400 font-medium mt-0.5">
			Pantau performa dan ringkasan tagihan Anda bulan ini
		</p>
	</div>
	<div class="flex items-center gap-4">
		<div
			class="bg-slate-100/80 p-1 rounded-2xl flex items-center border border-slate-200/60"
		>
			<button
				class="px-5 py-2 bg-white shadow-sm rounded-xl text-xs font-bold text-slate-900 transition-all"
				>View Admin</button
			>
			<button
				class="px-5 py-2 text-xs font-bold text-slate-400 hover:text-slate-700 transition-colors"
				>View Public</button
			>
		</div>
	</div>
</header>

<!-- Konten Utama Dashboard -->
<div class="p-10 space-y-8 max-w-7xl mx-auto w-full">
	<!-- Kartu Statistik -->
	<div class="grid grid-cols-1 md:grid-cols-3 gap-6">
		<div
			class="bg-white p-7 rounded-3xl shadow-sm border border-slate-100 hover:shadow-md transition-all duration-300"
		>
			<div class="flex justify-between items-center mb-4">
				<span class="text-xs font-bold text-slate-400 uppercase tracking-wider"
					>Total Pendapatan</span
				>
				<span
					class="inline-flex items-center gap-1 text-[11px] font-extrabold text-emerald-600 bg-emerald-50 px-2.5 py-1 rounded-full border border-emerald-100"
					>+12.5%</span
				>
			</div>
			<h3 class="text-3xl font-black text-slate-900 tracking-tight">
				{formatRupiah(stats.total_pendapatan)}
			</h3>
		</div>

		<div
			class="bg-white p-7 rounded-3xl shadow-sm border border-slate-100 hover:shadow-md transition-all duration-300"
		>
			<div class="flex justify-between items-center mb-4">
				<span class="text-xs font-bold text-slate-400 uppercase tracking-wider"
					>Menunggu Pembayaran</span
				>
				<span
					class="text-[11px] font-extrabold text-slate-600 bg-slate-100 px-2.5 py-1 rounded-full border border-slate-200/60"
					>{stats.menunggu_pembayaran.count} Invoice</span
				>
			</div>
			<h3 class="text-3xl font-black text-slate-900 tracking-tight">
				{formatRupiah(stats.menunggu_pembayaran.total)}
			</h3>
		</div>

		<div
			class="bg-white p-7 rounded-3xl shadow-sm border border-slate-100 hover:shadow-md transition-all duration-300"
		>
			<div class="flex justify-between items-center mb-4">
				<span class="text-xs font-bold text-slate-400 uppercase tracking-wider"
					>Jatuh Tempo</span
				>
				<span
					class="text-[11px] font-extrabold text-rose-600 bg-rose-50 px-2.5 py-1 rounded-full border border-rose-100"
					>{stats.jatuh_tempo.count} Invoice</span
				>
			</div>
			<h3 class="text-3xl font-black text-slate-900 tracking-tight">
				{formatRupiah(stats.jatuh_tempo.total)}
			</h3>
		</div>
	</div>

	<!-- Tabel Tagihan Terbaru -->
	<div
		class="bg-white rounded-3xl shadow-sm border border-slate-100 overflow-hidden"
	>
		<div
			class="px-8 py-7 border-b border-slate-100 flex items-center justify-between bg-white"
		>
			<div>
				<h3 class="text-lg font-black text-slate-900 tracking-tight">
					Tagihan Terbaru
				</h3>
				<p class="text-xs text-slate-400 font-medium mt-0.5">
					Daftar tagihan terintegrasi database
				</p>
			</div>
			<button
				onclick={() => (showModalLocal = true)}
				class="flex items-center gap-2 bg-slate-900 text-white px-5 py-3 rounded-2xl text-xs font-bold hover:bg-slate-800 transition-all shadow-lg shadow-slate-900/10 hover:-translate-y-0.5"
			>
				<svg
					class="w-4 h-4"
					fill="none"
					stroke="currentColor"
					viewBox="0 0 24 24"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2.5"
						d="M12 4v16m8-8H4"
					></path></svg
				>
				Buat Tagihan Baru
			</button>
		</div>

		<div class="overflow-x-auto">
			<table class="w-full text-left border-collapse">
				<thead>
					<tr
						class="bg-slate-50/60 text-slate-400 text-[11px] font-bold uppercase tracking-wider"
					>
						<th class="px-8 py-4">Nomor Invoice</th>
						<th class="px-8 py-4">Klien</th>
						<th class="px-8 py-4">Tanggal</th>
						<th class="px-8 py-4">Status</th>
						<th class="px-8 py-4 text-right">Jumlah</th>
						<th class="px-6 py-4"></th>
					</tr>
				</thead>
				<tbody class="divide-y divide-slate-100/80 text-sm">
					{#if isLoading}
						<tr>
							<td colspan="6" class="px-8 py-6 text-center text-slate-400 font-medium">
								Memuat data tagihan...
							</td>
						</tr>
					{:else if invoices.length === 0}
						<tr>
							<td colspan="6" class="px-8 py-6 text-center text-slate-400 font-medium">
								Belum ada tagihan yang dibuat.
							</td>
						</tr>
					{:else}
						{#each invoices as invoice}
							{@const statusClass = getStatusStyle(invoice.status)}
							<tr class="hover:bg-slate-50/50 transition-colors group">
								<td class="px-8 py-5 font-bold text-slate-900">{invoice.nomor_invoice}</td>
								<td class="px-8 py-5 font-semibold text-slate-600">
									{invoice.client ? invoice.client.nama : '-'}
								</td>
								<td class="px-8 py-5 text-slate-400 text-xs font-medium">
									{invoice.tanggal_terbit}
								</td>
								<td class="px-8 py-5">
									<span
										class="inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-xs font-bold {statusClass}"
									>
										<span class="w-1.5 h-1.5 rounded-full bg-current"></span>
										{invoice.status}
									</span>
								</td>
								<td class="px-8 py-5 font-black text-slate-900 text-right">
									{formatRupiah(invoice.total_amount)}
								</td>
								<td class="px-6 py-5 text-right">
									<button
										class="p-2 text-slate-300 hover:text-slate-700 rounded-xl hover:bg-slate-100 opacity-0 group-hover:opacity-100 transition-all"
									>
										<svg
											class="w-4 h-4"
											fill="none"
											stroke="currentColor"
											viewBox="0 0 24 24"
											><path
												stroke-linecap="round"
												stroke-linejoin="round"
												stroke-width="2.5"
												d="M12 6v.01M12 12v.01M12 18v.01"
											></path></svg
										>
									</button>
								</td>
							</tr>
						{/each}
					{/if}
				</tbody>
			</table>
		</div>
	</div>
</div>

<!-- Modal / Popup Overlay -->
{#if showModalLocal}
	<div
		class="fixed inset-0 z-[100] flex items-center justify-center p-4 sm:p-6 font-sans"
	>
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			class="absolute inset-0 bg-slate-900/40 backdrop-blur-sm transition-opacity"
			onclick={() => (showModalLocal = false)}
		></div>

		<div
			class="relative w-full max-w-2xl bg-white rounded-3xl shadow-2xl border border-slate-100 overflow-hidden transform transition-all flex flex-col max-h-[90vh]"
		>
			<div
				class="px-8 py-6 border-b border-slate-100 flex items-center justify-between shrink-0 bg-white"
			>
				<div>
					<h3 class="text-xl font-black text-slate-900 tracking-tight">
						Buat Tagihan Baru
					</h3>
					<p class="text-xs text-slate-400 font-medium mt-0.5">
						Masukkan detail informasi tagihan untuk klien Anda.
					</p>
				</div>
				<button
					onclick={() => (showModalLocal = false)}
					class="text-slate-400 hover:text-slate-700 transition-colors p-2 rounded-2xl hover:bg-slate-100"
				>
					<svg
						class="w-5 h-5"
						fill="none"
						stroke="currentColor"
						viewBox="0 0 24 24"
						><path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2.5"
							d="M6 18L18 6M6 6l12 12"
						></path></svg
					>
				</button>
			</div>

			<div class="px-8 py-6 overflow-y-auto space-y-6">
				<form onsubmit={handleSubmitLocal} class="space-y-5" id="invoiceForm">
					<div class="grid grid-cols-1 md:grid-cols-2 gap-5">
						<div>
							<label
								class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
								>Nomor Invoice (Opsional)</label
							>
							<input
								type="text"
								bind:value={nomorInvoice}
								placeholder="Otomatis jika kosong"
								class="w-full px-4 py-3 rounded-2xl bg-slate-50 border border-slate-200/80 text-slate-900 text-sm font-bold focus:outline-none"
							/>
						</div>
						<div>
							<label
								class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
								>Tanggal Terbit</label
							>
							<input
								type="date"
								bind:value={tanggalTerbit}
								required
								class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-800 text-sm font-medium focus:ring-2 focus:ring-slate-900 focus:border-transparent outline-none transition-all shadow-sm"
							/>
						</div>
					</div>

					<div class="grid grid-cols-1 md:grid-cols-2 gap-5">
						<div>
							<label
								class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
								>Pilih Klien</label
							>
							<select
								bind:value={selectedClientId}
								required
								class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-800 text-sm font-semibold focus:ring-2 focus:ring-slate-900 focus:border-transparent outline-none transition-all cursor-pointer shadow-sm"
							>
								<option value="" disabled selected>Pilih klien dari daftar...</option>
								{#each clients as client}
									<option value={client.ID}>{client.nama}</option>
								{/each}
							</select>
						</div>
						<div>
							<label
								class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
								>Tanggal Jatuh Tempo</label
							>
							<input
								type="date"
								bind:value={tanggalJatuhTempo}
								required
								class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-800 text-sm font-medium focus:ring-2 focus:ring-slate-900 focus:border-transparent outline-none transition-all shadow-sm"
							/>
						</div>
					</div>

					<div>
						<label
							class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
							>Jumlah Tagihan (Rp)</label
						>
						<div class="relative">
							<div
								class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none"
							>
								<span class="text-slate-400 font-bold text-sm">Rp</span>
							</div>
							<input
								type="number"
								bind:value={jumlahTagihan}
								placeholder="0"
								min="0"
								required
								class="w-full pl-12 pr-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-900 text-sm font-bold focus:ring-2 focus:ring-slate-900 focus:border-transparent outline-none transition-all shadow-sm"
							/>
						</div>
					</div>

					<div>
						<label
							class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
							>Deskripsi Layanan / Catatan</label
						>
						<textarea
							rows="3"
							bind:value={deskripsi}
							placeholder="Misal: Jasa Pembuatan Website Company Profile"
							class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-800 text-sm font-medium focus:ring-2 focus:ring-slate-900 focus:border-transparent outline-none transition-all resize-none shadow-sm"
						></textarea>
					</div>
				</form>
			</div>

			<div
				class="px-8 py-5 border-t border-slate-100 flex items-center justify-end gap-3 bg-slate-50/50 shrink-0"
			>
				<button
					type="button"
					onclick={() => (showModalLocal = false)}
					class="px-5 py-2.5 rounded-2xl text-xs font-bold text-slate-500 hover:text-slate-900 hover:bg-slate-200/50 transition-colors"
					>Batal</button
				>
				<button
					type="submit"
					form="invoiceForm"
					disabled={isSubmitting}
					class="px-6 py-2.5 rounded-2xl bg-slate-900 text-white text-xs font-bold hover:bg-slate-800 transition-all shadow-lg shadow-slate-900/10 hover:-translate-y-0.5 disabled:opacity-50"
					>{isSubmitting ? 'Memproses...' : 'Simpan Tagihan'}</button
				>
			</div>
		</div>
	</div>
{/if}