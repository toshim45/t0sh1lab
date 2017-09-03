<?php

use Carbon\Carbon;
use Illuminate\Database\Seeder;

class ItemTableSeeder extends Seeder {

    public function run()
    {
        $timeStampNow = Carbon::now()->toDateTimeString();
        DB::table('items')->insert(['id'=>1, 'name'=>'item-1', 'amount'=>1, 'active'=>1, 'created_at'=>$timeStampNow, 'updated_at'=>$timeStampNow]);
    }

}