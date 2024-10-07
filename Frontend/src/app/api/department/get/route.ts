import { NextResponse } from 'next/server';
import { $serverReq } from '@/utils/serverRequest';

export async function GET() {
  try {
    const data = await $serverReq.get('/api/department');
    return NextResponse.json(data);
  }  catch (error: any) {
    return NextResponse.json(
      { error: error.message },
      { status: error.status || 500 }
    );
  }
}