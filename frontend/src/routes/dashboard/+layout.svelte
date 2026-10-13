<script>
    import { page } from "$app/stores";
    let { children } = $props();
    let showModal = $state(false);
</script>

<div class="flex h-screen bg-slate-50 font-sans text-slate-900 overflow-hidden">
    <!-- Sidebar -->
    <aside
        class="w-64 bg-white border-r border-slate-100 flex flex-col justify-between z-20 shrink-0"
    >
        <div>
            <div class="h-24 flex items-center px-8">
                <div
                    class="bg-slate-900 text-white p-2.5 rounded-2xl mr-3 shadow-md shadow-slate-900/10 flex items-center justify-center shrink-0 w-10 h-10"
                >
                    <svg
                        class="w-5 h-5"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2.5"
                            d="M13 10V3L4 14h7v7l9-11h-7z"
                        />
                    </svg>
                </div>
                <span class="text-xl font-black tracking-tight text-slate-900"
                    >Kilat</span
                >
            </div>

            <nav class="px-4 py-2 space-y-1">
                <a
                    href="/dashboard"
                    class="flex items-center px-4 py-3.5 rounded-2xl font-bold text-xs transition-colors {$page
                        .url.pathname === '/dashboard'
                        ? 'bg-slate-900 text-white shadow-lg shadow-slate-900/10'
                        : 'text-slate-500 hover:bg-slate-50 hover:text-slate-900'}"
                >
                    <svg
                        class="w-4 h-4 mr-3 shrink-0"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        ><path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2.5"
                            d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zm10 0a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zm10 0a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"
                        /></svg
                    >
                    Dashboard
                </a>

                <button
                    onclick={() => (showModal = true)}
                    class="w-full flex items-center px-4 py-3.5 rounded-2xl font-bold text-xs text-slate-500 hover:bg-slate-50 hover:text-slate-900 transition-colors text-left"
                >
                    <svg
                        class="w-4 h-4 mr-3 shrink-0"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        ><path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2.5"
                            d="M12 4v16m8-8H4"
                        /></svg
                    >
                    Buat Tagihan
                </button>

                <a
                    href="/dashboard/klien"
                    class="flex items-center px-4 py-3.5 rounded-2xl font-bold text-xs transition-colors {$page.url.pathname.includes(
                        '/klien',
                    )
                        ? 'bg-slate-900 text-white shadow-lg shadow-slate-900/10'
                        : 'text-slate-500 hover:bg-slate-50 hover:text-slate-900'}"
                >
                    <svg
                        class="w-4 h-4 mr-3 shrink-0"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        ><path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2.5"
                            d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
                        /></svg
                    >
                    Klien
                </a>

                <!-- Menu Pengaturan -->
                <a
                    href="/dashboard/pengaturan"
                    class="flex items-center px-4 py-3.5 rounded-2xl font-bold text-xs transition-colors {$page.url.pathname.includes(
                        '/pengaturan',
                    )
                        ? 'bg-slate-900 text-white shadow-lg shadow-slate-900/10'
                        : 'text-slate-500 hover:bg-slate-50 hover:text-slate-900'}"
                >
                    <svg
                        class="w-4 h-4 mr-3 shrink-0"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        ><path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2.5"
                            d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 2.51.55 1.566 1.49-.691.691-.94 1.756-.35 2.573.94 1.543-.55 2.51-1.49 1.566-.691-.691-1.756-.94-2.573-.35-1.756.426-1.756 2.924 0 3.35 1.066.257 1.724 1.066 2.573 1.543.94-.55 2.51-.94 1.566-1.49-.691-.691-.94-1.756-.35-2.573.94-1.543-.55-2.51-1.49-1.566zM15 12a3 3 0 11-6 0 3 3 0 016 0z"
                        /></svg
                    >
                    Pengaturan
                </a>
            </nav>
        </div>

        <div class="p-4 m-4 bg-slate-50 rounded-3xl border border-slate-100">
            <p class="text-xs font-black text-slate-800 mb-1">Paket Pro</p>
            <p class="text-[10px] text-slate-400 font-bold mb-2.5">
                Kuota 120 invoice/bulan
            </p>
            <div
                class="w-full bg-slate-200 h-1.5 rounded-full overflow-hidden mb-2"
            >
                <div class="bg-slate-900 h-full w-[66%]"></div>
            </div>
            <p class="text-[10px] text-slate-400 font-medium">
                80 dari 120 terpakai
            </p>
        </div>
    </aside>

    <!-- Area Konten Utama -->
    <main class="flex-1 flex flex-col overflow-y-auto relative bg-transparent">
        {@render children()}
    </main>
</div>

<!-- Modal Global Buat Tagihan -->
{#if showModal}
    <div
        class="fixed inset-0 z-[100] flex items-center justify-center p-4 sm:p-6 font-sans"
    >
        <div
            class="absolute inset-0 bg-slate-900/40 backdrop-blur-sm transition-opacity"
            onclick={() => (showModal = false)}
        ></div>
        <div
            class="relative w-full max-w-2xl bg-white rounded-3xl shadow-2xl border border-slate-100 overflow-hidden transform transition-all flex flex-col max-h-[90vh]"
        >
            <div
                class="px-8 py-6 border-b border-slate-100 flex items-center justify-between shrink-0 bg-white"
            >
                <div>
                    <h3
                        class="text-xl font-black text-slate-900 tracking-tight"
                    >
                        Buat Tagihan Baru
                    </h3>
                    <p class="text-xs text-slate-400 font-medium mt-0.5">
                        Masukkan detail informasi tagihan untuk klien.
                    </p>
                </div>
                <button
                    onclick={() => (showModal = false)}
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
                        /></svg
                    >
                </button>
            </div>

            <div class="px-8 py-6 overflow-y-auto">
                <form class="space-y-5">
                    <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                        <div>
                            <label
                                class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                                >Nomor Invoice</label
                            >
                            <input
                                type="text"
                                value="INV-2026-092"
                                class="w-full px-4 py-3 rounded-2xl bg-slate-50 border border-slate-200/80 text-slate-900 text-sm font-bold focus:outline-none"
                                readonly
                            />
                        </div>
                        <div>
                            <label
                                class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                                >Pilih Klien</label
                            >
                            <select
                                class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-900 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-slate-900 focus:border-transparent cursor-pointer appearance-none"
                            >
                                <option value="" disabled selected
                                    >Pilih klien dari daftar...</option
                                >
                                <option value="1">PT Sinar Abadi</option>
                                <option value="2">CV Mitra Jaya</option>
                                <option value="3">Studio Nusantara</option>
                                <option value="4">Toko Berkah</option>
                            </select>
                        </div>
                    </div>

                    <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
                        <div>
                            <label
                                class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                                >Tanggal Terbit</label
                            >
                            <input
                                type="date"
                                class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-900 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-slate-900 focus:border-transparent"
                            />
                        </div>
                        <div>
                            <label
                                class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                                >Tanggal Jatuh Tempo</label
                            >
                            <input
                                type="date"
                                class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-900 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-slate-900 focus:border-transparent"
                            />
                        </div>
                    </div>

                    <div>
                        <label
                            class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                            >Deskripsi Layanan / Produk</label
                        >
                        <input
                            type="text"
                            placeholder="Misal: Jasa Pembuatan Website Company Profile"
                            class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-900 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-slate-900 focus:border-transparent"
                        />
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
                                <span class="text-slate-400 font-bold text-sm"
                                    >Rp</span
                                >
                            </div>
                            <input
                                type="number"
                                placeholder="0"
                                class="w-full pl-12 pr-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-900 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-slate-900 focus:border-transparent"
                            />
                        </div>
                    </div>
                </form>
            </div>

            <div
                class="px-8 py-5 border-t border-slate-100 bg-slate-50 flex justify-end gap-3 shrink-0"
            >
                <button
                    onclick={() => (showModal = false)}
                    class="px-6 py-2.5 rounded-xl font-bold text-sm text-slate-600 bg-white border border-slate-200 hover:bg-slate-50 transition-colors"
                >
                    Batal
                </button>
                <button
                    onclick={() => {
                        alert("Tagihan baru berhasil dibuat!");
                        showModal = false;
                    }}
                    class="px-6 py-2.5 rounded-xl font-bold text-sm text-white bg-slate-900 hover:bg-slate-800 shadow-lg shadow-slate-900/20 transition-all hover:-translate-y-0.5"
                >
                    Simpan Tagihan
                </button>
            </div>
        </div>
    </div>
{/if}
