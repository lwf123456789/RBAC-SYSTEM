import { NextResponse } from 'next/server';
import { $serverReq } from '@/utils/serverRequest';

export async function GET() {
  try {
    const data = await $serverReq.get('/api/dicts?type=ROLE_TYPE');
    return NextResponse.json(data);
  }  catch (error: any) {
    return NextResponse.json(
      { error: error.message },
      { status: error.status || 500 }
    );
  }
}