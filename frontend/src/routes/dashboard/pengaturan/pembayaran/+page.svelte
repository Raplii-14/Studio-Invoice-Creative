<script>
    // State daftar rekening pembayaran
    let listRekening = $state([
        {
            id: 1,
            bank: "Bank BCA",
            nomor: "1234 5678 9012",
            pemilik: "Kilat Studio",
            aktif: true,
        },
        {
            id: 2,
            bank: "Bank Mandiri",
            nomor: "9876 5432 1098",
            pemilik: "Kilat Studio",
            aktif: true,
        },
    ]);

    // State untuk modal tambah rekening
    let showAddModal = $state(false);
    let namaBank = $state("");
    let nomorRekening = $state("");
    let namaPemilik = $state("");

    function handleTambah(e) {
        e.preventDefault();
        if (!namaBank || !nomorRekening || !namaPemilik) return;

        const rekeningBaru = {
            id: Date.now(),
            bank: namaBank,
            nomor: nomorRekening,
            pemilik: namaPemilik,
            aktif: true,
        };

        listRekening = [...listRekening, rekeningBaru];

        // Reset form & tutup modal
        namaBank = "";
        nomorRekening = "";
        namaPemilik = "";
        showAddModal = false;
    }

    function hapusRekening(id) {
        listRekening = listRekening.filter((item) => item.id !== id);
    }
</script>

<div class="space-y-6 max-w-3xl">
    <!-- Header Bagian -->
    <div
        class="border-b border-slate-100 pb-4 flex items-center justify-between"
    >
        <div>
            <h3 class="text-xl font-black text-slate-900 tracking-tight">
                Metode Pembayaran
            </h3>
            <p class="text-xs text-slate-400 font-medium mt-0.5">
                Rekening yang tampil di tagihan klien Anda
            </p>
        </div>
        <button
            onclick={() => (showAddModal = true)}
            class="px-4 py-2.5 rounded-2xl font-bold text-xs text-white bg-slate-900 hover:bg-slate-800 shadow-md shadow-slate-950/10 transition-all hover:-translate-y-0.5 flex items-center gap-2"
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
                />
            </svg>
            Tambah Rekening
        </button>
    </div>

    <!-- Daftar Rekening -->
    <div class="space-y-4">
        {#each listRekening as rek (rek.id)}
            <div
                class="bg-white rounded-3xl border border-slate-200/80 p-6 flex items-center justify-between shadow-sm transition-all hover:shadow-md"
            >
                <div class="flex items-center gap-4">
                    <div
                        class="w-12 h-12 rounded-2xl bg-slate-50 border border-slate-200/80 flex items-center justify-center text-slate-700 shrink-0"
                    >
                        <svg
                            class="w-6 h-6"
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                        >
                            <path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="2"
                                d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"
                            />
                        </svg>
                    </div>
                    <div>
                        <div class="flex items-center gap-2">
                            <p class="text-sm font-black text-slate-900">
                                {rek.bank}
                            </p>
                            <span
                                class="inline-flex items-center px-2 py-0.5 rounded-full text-[10px] font-bold bg-emerald-50 text-emerald-600 border border-emerald-200/60"
                                >Aktif</span
                            >
                        </div>
                        <p class="text-xs font-semibold text-slate-600 mt-1">
                            {rek.nomor} &bull; a.n. {rek.pemilik}
                        </p>
                    </div>
                </div>

                <button
                    onclick={() => hapusRekening(rek.id)}
                    class="text-slate-400 hover:text-rose-600 p-2 rounded-xl hover:bg-rose-50 transition-colors"
                    title="Hapus Rekening"
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
                            stroke-width="2"
                            d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                        />
                    </svg>
                </button>
            </div>
        {/each}

        {#if listRekening.length === 0}
            <div
                class="text-center py-12 bg-white rounded-3xl border border-slate-200/80"
            >
                <p class="text-xs text-slate-400 font-medium">
                    Belum ada rekening pembayaran yang ditambahkan.
                </p>
            </div>
        {/if}
    </div>
</div>

<!-- Modal / Popup Tambah Rekening -->
{#if showAddModal}
    <div
        class="fixed inset-0 z-[100] flex items-center justify-center p-4 sm:p-6 font-sans"
    >
        <!-- Overlay -->
        <!-- svelte-ignore a11y_click_events_have_key_events -->
        <!-- svelte-ignore a11y_no_static_element_interactions -->
        <div
            class="absolute inset-0 bg-slate-900/60 backdrop-blur-sm transition-opacity opacity-100"
            onclick={() => (showAddModal = false)}
        ></div>

        <!-- Konten Modal -->
        <div
            class="relative w-full max-w-lg bg-white rounded-3xl shadow-2xl overflow-hidden transform transition-all duration-300 scale-100 flex flex-col"
        >
            <div
                class="px-8 py-6 border-b border-slate-100 flex items-center justify-between shrink-0 bg-white"
            >
                <div>
                    <h3
                        class="text-xl font-black text-slate-900 tracking-tight"
                    >
                        Tambah Rekening Baru
                    </h3>
                    <p class="text-xs text-slate-400 font-medium mt-0.5">
                        Masukkan detail rekening bank untuk tujuan pembayaran.
                    </p>
                </div>
                <button
                    onclick={() => (showAddModal = false)}
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

            <form onsubmit={handleTambah} class="px-8 py-6 space-y-4">
                <div>
                    <label
                        class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                        >Nama Bank / E-Wallet</label
                    >
                    <input
                        type="text"
                        bind:value={namaBank}
                        placeholder="Misal: Bank BNI / GoPay"
                        required
                        class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-900 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-slate-900 transition-all shadow-sm"
                    />
                </div>

                <div>
                    <label
                        class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                        >Nomor Rekening</label
                    >
                    <input
                        type="text"
                        bind:value={nomorRekening}
                        placeholder="Misal: 1234567890"
                        required
                        class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-900 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-slate-900 transition-all shadow-sm"
                    />
                </div>

                <div>
                    <label
                        class="block text-xs font-bold text-slate-700 uppercase tracking-wider mb-2"
                        >Atas Nama (Pemilik Rekening)</label
                    >
                    <input
                        type="text"
                        bind:value={namaPemilik}
                        placeholder="Misal: Kilat Studio / Budi"
                        required
                        class="w-full px-4 py-3 rounded-2xl bg-white border border-slate-200/80 text-slate-900 text-sm font-medium focus:outline-none focus:ring-2 focus:ring-slate-900 transition-all shadow-sm"
                    />
                </div>

                <div
                    class="pt-4 border-t border-slate-100 flex justify-end gap-3"
                >
                    <button
                        type="button"
                        onclick={() => (showAddModal = false)}
                        class="px-6 py-2.5 rounded-xl font-bold text-sm text-slate-600 bg-white border border-slate-200 hover:bg-slate-50 transition-colors"
                    >
                        Batal
                    </button>
                    <button
                        type="submit"
                        class="px-6 py-2.5 rounded-xl font-bold text-sm text-white bg-slate-900 hover:bg-slate-800 shadow-lg shadow-slate-900/20 transition-all hover:-translate-y-0.5"
                    >
                        Simpan Rekening
                    </button>
                </div>
            </form>
        </div>
    </div>
{/if}
