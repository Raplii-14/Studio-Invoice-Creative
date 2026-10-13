<script>
    let showAddClientModal = $state(false);

    let clients = [
        {
            initials: "P",
            name: "PT Sinar Abadi",
            email: "finance@sinarabadi.co.id",
            phone: "081234567890",
            address: "Jl. Pemuda No. 12, Surabaya",
            total: "Rp 124.500.000",
            invoices: 12,
            status: "Lunas",
        },
        {
            initials: "C",
            name: "CV Mitra Jaya",
            email: "admin@mitrajaya.id",
            phone: "082198765432",
            address: "Jl. Raya Darmo No. 45, Surabaya",
            total: "Rp 45.200.000",
            invoices: 8,
            status: "Belum Bayar",
        },
        {
            initials: "S",
            name: "Studio Nusantara",
            email: "halo@studionusantara.com",
            phone: "085611223344",
            address: "Jl. Tunjungan No. 88, Surabaya",
            total: "Rp 18.900.000",
            invoices: 5,
            status: "Jatuh Tempo",
        },
        {
            initials: "T",
            name: "Toko Berkah",
            email: "tokoberkah@gmail.com",
            phone: "087855443322",
            address: "Jl. Pahlawan No. 10, Sidoarjo",
            total: "Rp 67.300.000",
            invoices: 15,
            status: "Lunas",
        },
        {
            initials: "A",
            name: "Andi Pratama",
            email: "andi.pratama@gmail.com",
            phone: "081399887766",
            address: "Perum Graha Asri Blok C-5, Gresik",
            total: "Rp 9.600.000",
            invoices: 3,
            status: "Belum Bayar",
        },
        {
            initials: "L",
            name: "Lestari Group",
            email: "ap@lestarigroup.co.id",
            phone: "081933344455",
            address: "Jl. HR Mohammad No. 21, Surabaya",
            total: "Rp 210.400.000",
            invoices: 20,
            status: "Lunas",
        },
    ];

    /** @param {string} status */
    function getStatusBadge(status) {
        if (status === "Lunas")
            return "bg-emerald-50 text-emerald-600 border border-emerald-100";
        if (status === "Belum Bayar")
            return "bg-amber-50 text-amber-600 border border-amber-100";
        return "bg-rose-50 text-rose-600 border border-rose-100";
    }

    /** @param {SubmitEvent} event */
    function handleAddClient(event) {
        event.preventDefault();
        alert("Klien baru berhasil ditambahkan!");
        showAddClientModal = false;
    }
</script>

<!-- Header Halaman Klien -->
<header
    class="h-24 flex items-center justify-between px-10 bg-white/70 backdrop-blur-md sticky top-0 z-10 border-b border-slate-100/80"
>
    <div>
        <h1 class="text-2xl font-black text-slate-900 tracking-tight">Klien</h1>
        <p class="text-sm text-slate-400 font-medium mt-0.5">
            Kelola daftar perusahaan dan mitra klien Anda
        </p>
    </div>
</header>

<div class="p-10 space-y-8 max-w-7xl mx-auto w-full">
    <!-- Action Bar: Search & Tambah -->
    <div
        class="flex flex-col sm:flex-row justify-between items-center gap-4 bg-white p-4 rounded-3xl border border-slate-100 shadow-sm"
    >
        <div class="relative w-full sm:w-96">
            <svg
                class="w-4 h-4 absolute left-4 top-3.5 text-slate-400"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
            >
                <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2.5"
                    d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                ></path>
            </svg>
            <input
                type="text"
                placeholder="Cari nama atau email klien..."
                class="w-full pl-11 pr-4 py-2.5 rounded-2xl bg-slate-50 border border-slate-200/80 text-sm font-medium text-slate-800 focus:outline-none focus:bg-white focus:border-slate-900 transition-all"
            />
        </div>
        <button
            onclick={() => (showAddClientModal = true)}
            class="w-full sm:w-auto bg-slate-900 text-white px-6 py-3 rounded-2xl text-xs font-bold flex items-center justify-center gap-2 hover:bg-slate-800 transition-all shadow-lg shadow-slate-900/10 hover:-translate-y-0.5"
        >
            <svg
                class="w-4 h-4"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
            >
                <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2.5"
                    d="M12 4v16m8-8H4"
                ></path>
            </svg>
            Tambah Klien
        </button>
    </div>

    <!-- Grid Kartu Klien -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {#each clients as client}
            <div
                class="bg-white rounded-3xl p-7 shadow-sm border border-slate-100 flex flex-col justify-between hover:shadow-md transition-all duration-300 group"
            >
                <div>
                    <!-- Top Section: Avatar & Info -->
                    <div class="flex items-start justify-between mb-4">
                        <div class="flex items-center gap-4">
                            <div
                                class="w-12 h-12 rounded-2xl bg-slate-900 text-white flex items-center justify-center font-black text-lg shadow-md shadow-slate-900/10 shrink-0"
                            >
                                {client.initials}
                            </div>
                            <div class="overflow-hidden">
                                <h3
                                    class="font-bold text-slate-900 truncate tracking-tight"
                                >
                                    {client.name}
                                </h3>
                                <p
                                    class="text-xs text-slate-400 font-medium truncate mt-0.5"
                                >
                                    {client.email}
                                </p>
                            </div>
                        </div>
                        <button
                            class="text-slate-300 hover:text-slate-600 p-1.5 rounded-xl hover:bg-slate-50 transition-colors"
                        >
                            <svg
                                class="w-5 h-5"
                                fill="currentColor"
                                viewBox="0 0 24 24"
                            >
                                <path
                                    d="M12 8c1.1 0 2-.9 2-2s-.9-2-2-2-2 .9-2 2 .9 2 2 2zm0 2c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"
                                ></path>
                            </svg>
                        </button>
                    </div>

                    <!-- Info Tambahan (Telepon & Alamat) -->
                    <div
                        class="space-y-1.5 mb-5 text-xs text-slate-500 font-medium bg-slate-50/70 p-3.5 rounded-2xl border border-slate-100"
                    >
                        <p class="flex items-center gap-2 truncate">
                            <svg
                                class="w-3.5 h-3.5 text-slate-400 shrink-0"
                                fill="none"
                                stroke="currentColor"
                                viewBox="0 0 24 24"
                                ><path
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    stroke-width="2"
                                    d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"
                                /></svg
                            >
                            {client.phone}
                        </p>
                        <p class="flex items-center gap-2 truncate">
                            <svg
                                class="w-3.5 h-3.5 text-slate-400 shrink-0"
                                fill="none"
                                stroke="currentColor"
                                viewBox="0 0 24 24"
                                ><path
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    stroke-width="2"
                                    d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"
                                /><path
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    stroke-width="2"
                                    d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"
                                /></svg
                            >
                            {client.address}
                        </p>
                    </div>
                </div>

                <!-- Bottom Section: Stats -->
                <div
                    class="flex justify-between items-end pt-3 border-t border-slate-100"
                >
                    <div>
                        <p
                            class="text-[10px] text-slate-400 uppercase tracking-widest font-bold mb-1"
                        >
                            Total Tagihan
                        </p>
                        <p class="font-black text-slate-900 tracking-tight">
                            {client.total}
                        </p>
                    </div>
                    <div class="text-right">
                        <p class="text-xs text-slate-400 font-medium mb-1.5">
                            {client.invoices} invoice
                        </p>
                        <span
                            class="inline-block px-2.5 py-1 rounded-full text-[11px] font-bold {getStatusBadge(
                                client.status,
                            )}"
                        >
                            {client.status}
                        </span>
                    </div>
                </div>
            </div>
        {/each}
    </div>
</div>

<!-- Modal Tambah Klien Lebih Detail -->
{#if showAddClientModal}
    <div
        class="fixed inset-0 z-[100] flex items-center justify-center p-4 sm:p-6 font-sans"
    >
        <!-- svelte-ignore a11y_click_events_have_key_events -->
        <!-- svelte-ignore a11y_no_static_element_interactions -->
        <div
            class="absolute inset-0 bg-slate-900/40 backdrop-blur-sm transition-opacity"
            onclick={() => (showAddClientModal = false)}
        ></div>

        <div
            class="relative w-full max-w-xl bg-white rounded-3xl shadow-2xl border border-slate-100 overflow-hidden transform transition-all flex flex-col max-h-[90vh]"
        >
            <div
                class="px-8 py-6 border-b border-slate-100 flex items-center justify-between shrink-0 bg-white"
            >
                <div>
                    <h3
                        class="text-xl font-black text-slate-900 tracking-tight"
                    >
                        Tambah Klien Baru
                    </h3>
                    <p class="text-xs text-slate-400 font-medium mt-0.5">
                        Masukkan informasi detail perusahaan atau perorangan
                        klien.
                    </p>
                </div>
                <button
                    onclick={() => (showAddClientModal = false)}
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

            <div class="px-8 py-6 overflow-y-auto space-y-4">
                <form
                    onsubmit={handleAddClient}
                    class="space-y-4"
                    id="clientForm"
                >
                    <div>
                        <label
                            class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                            >Nama Klien / Perusahaan</label
                        >
                        <input
                            type="text"
                            placeholder="Contoh: PT Sinar Abadi"
                            required
                            class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-800 text-sm font-medium focus:ring-2 focus:ring-slate-900 focus:border-transparent outline-none transition-all shadow-sm"
                        />
                    </div>
                    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                        <div>
                            <label
                                class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                                >Email Klien</label
                            >
                            <input
                                type="email"
                                placeholder="finance@sinarabadi.co.id"
                                required
                                class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-800 text-sm font-medium focus:ring-2 focus:ring-slate-900 focus:border-transparent outline-none transition-all shadow-sm"
                            />
                        </div>
                        <div>
                            <label
                                class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                                >Nomor Telepon / WhatsApp</label
                            >
                            <input
                                type="text"
                                placeholder="081234567890"
                                required
                                class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-800 text-sm font-medium focus:ring-2 focus:ring-slate-900 focus:border-transparent outline-none transition-all shadow-sm"
                            />
                        </div>
                    </div>
                    <div>
                        <label
                            class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                            >Alamat Lengkap</label
                        >
                        <textarea
                            rows="2"
                            placeholder="Jl. Pemuda No. 12, Surabaya"
                            required
                            class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-800 text-sm font-medium focus:ring-2 focus:ring-slate-900 focus:border-transparent outline-none transition-all resize-none shadow-sm"
                        ></textarea>
                    </div>
                    <div>
                        <label
                            class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                            >Catatan Tambahan (Opsional)</label
                        >
                        <textarea
                            rows="2"
                            placeholder="Catatan khusus terkait klien ini..."
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
                    onclick={() => (showAddClientModal = false)}
                    class="px-5 py-2.5 rounded-2xl text-xs font-bold text-slate-500 hover:text-slate-900 hover:bg-slate-200/50 transition-colors"
                    >Batal</button
                >
                <button
                    type="submit"
                    form="clientForm"
                    class="px-6 py-2.5 rounded-2xl bg-slate-900 text-white text-xs font-bold hover:bg-slate-800 transition-all shadow-lg shadow-slate-900/10 hover:-translate-y-0.5"
                    >Simpan Klien</button
                >
            </div>
        </div>
    </div>
{/if}
