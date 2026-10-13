<script>
    let autoReminder = $state(true);
    let weeklyReport = $state(false);
    let emailNotif = $state(true);

    function handleSave(e) {
        e.preventDefault();
        alert("Pengaturan notifikasi berhasil disimpan!");
    }
</script>

<div class="space-y-6 max-w-3xl">
    <!-- Header Bagian -->
    <div
        class="border-b border-slate-100 pb-4 flex items-center justify-between"
    >
        <div>
            <h3 class="text-xl font-black text-slate-900 tracking-tight">
                Notifikasi
            </h3>
            <p class="text-xs text-slate-400 font-medium mt-0.5">
                Atur bagaimana dan kapan Anda menerima pemberitahuan
            </p>
        </div>
        <button
            onclick={handleSave}
            class="px-4 py-2.5 rounded-2xl font-bold text-xs text-white bg-slate-900 hover:bg-slate-800 shadow-md shadow-slate-950/10 transition-all hover:-translate-y-0.5 flex items-center gap-2"
        >
            Simpan Perubahan
        </button>
    </div>

    <!-- Konten Pengaturan Notifikasi -->
    <div
        class="bg-white rounded-3xl border border-slate-200/80 p-8 space-y-6 shadow-sm"
    >
        <!-- Opsi 1: Pengingat Otomatis -->
        <div
            class="flex items-center justify-between pb-6 border-b border-slate-100"
        >
            <div class="space-y-0.5 pr-4">
                <p class="text-sm font-bold text-slate-900">
                    Pengingat Otomatis WhatsApp
                </p>
                <p class="text-xs text-slate-500 font-medium">
                    Kirim pengingat tagihan ke klien secara otomatis 2 hari
                    sebelum jatuh tempo.
                </p>
            </div>

            <!-- Toggle Switch Flexbox Murni -->
            <!-- svelte-ignore a11y_click_events_have_key_events -->
            <!-- svelte-ignore a11y_interactive_supports_focus -->
            <div
                role="switch"
                aria-checked={autoReminder}
                onclick={() => (autoReminder = !autoReminder)}
                class="flex w-14 h-8 items-center rounded-full p-1 cursor-pointer transition-colors duration-300 shrink-0 {autoReminder
                    ? 'bg-slate-900'
                    : 'bg-slate-200'}"
            >
                <div
                    class="w-6 h-6 bg-white rounded-full shadow-md transition-transform duration-300 ease-in-out"
                    style="transform: translateX({autoReminder
                        ? '24px'
                        : '0px'});"
                ></div>
            </div>
        </div>

        <!-- Opsi 2: Laporan Mingguan -->
        <div
            class="flex items-center justify-between pb-6 border-b border-slate-100"
        >
            <div class="space-y-0.5 pr-4">
                <p class="text-sm font-bold text-slate-900">
                    Laporan Pendapatan Mingguan
                </p>
                <p class="text-xs text-slate-500 font-medium">
                    Terima ringkasan total pendapatan dan statistik tagihan
                    setiap Senin pagi.
                </p>
            </div>

            <!-- svelte-ignore a11y_click_events_have_key_events -->
            <!-- svelte-ignore a11y_interactive_supports_focus -->
            <div
                role="switch"
                aria-checked={weeklyReport}
                onclick={() => (weeklyReport = !weeklyReport)}
                class="flex w-14 h-8 items-center rounded-full p-1 cursor-pointer transition-colors duration-300 shrink-0 {weeklyReport
                    ? 'bg-slate-900'
                    : 'bg-slate-200'}"
            >
                <div
                    class="w-6 h-6 bg-white rounded-full shadow-md transition-transform duration-300 ease-in-out"
                    style="transform: translateX({weeklyReport
                        ? '24px'
                        : '0px'});"
                ></div>
            </div>
        </div>

        <!-- Opsi 3: Notifikasi Email -->
        <div class="flex items-center justify-between">
            <div class="space-y-0.5 pr-4">
                <p class="text-sm font-bold text-slate-900">
                    Notifikasi Pembayaran Masuk
                </p>
                <p class="text-xs text-slate-500 font-medium">
                    Kirim email instan saat klien berhasil melunasi tagihan
                    mereka.
                </p>
            </div>

            <!-- svelte-ignore a11y_click_events_have_key_events -->
            <!-- svelte-ignore a11y_interactive_supports_focus -->
            <div
                role="switch"
                aria-checked={emailNotif}
                onclick={() => (emailNotif = !emailNotif)}
                class="flex w-14 h-8 items-center rounded-full p-1 cursor-pointer transition-colors duration-300 shrink-0 {emailNotif
                    ? 'bg-slate-900'
                    : 'bg-slate-200'}"
            >
                <div
                    class="w-6 h-6 bg-white rounded-full shadow-md transition-transform duration-300 ease-in-out"
                    style="transform: translateX({emailNotif
                        ? '24px'
                        : '0px'});"
                ></div>
            </div>
        </div>
    </div>
</div>
